/*
 * (c) Copyright IBM Corp. 2021
 * (c) Copyright Instana Inc. 2021
 */

package controllers

import (
	"context"
	"errors"

	"github.com/go-logr/logr"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"

	instanav1 "github.com/instana/instana-agent-operator/api/v1"
	"github.com/instana/instana-agent-operator/controllers/reconciliation/helm"
	instanaclient "github.com/instana/instana-agent-operator/pkg/k8s/client"
)

const (
	finalizerV1 = "agent.instana.io/finalizer"
	finalizerV3 = "agent.instana.io/finalizer/v3"
)

// Add will create a new Instana Agent Controller and add this to the Manager for reconciling
func Add(mgr manager.Manager) error {
	return add(
		mgr, NewInstanaAgentReconciler(
			mgr.GetClient(),
			mgr.GetScheme(),
			mgr.GetEventRecorderFor("agent.controller"),
			logf.Log.WithName("agent.controller"),
		),
	)
}

// add sets up the controller with the Manager.
func add(mgr ctrl.Manager, r *InstanaAgentReconciler) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&instanav1.InstanaAgent{}).
		// TODO: Update list of Owns
		Owns(&appsv1.DaemonSet{}).
		Owns(&corev1.Pod{}).
		Owns(&corev1.ConfigMap{}).
		Owns(&corev1.Secret{}).
		Owns(&corev1.ServiceAccount{}).
		Owns(&corev1.Service{}).
		Owns(&rbacv1.ClusterRole{}).
		Owns(&rbacv1.ClusterRoleBinding{}).
		WithEventFilter(filterPredicate()).
		Complete(r)
}

// Create generic filter for all events, that removes some chattiness mainly when only the Status field has been updated.
func filterPredicate() predicate.Predicate {
	return predicate.Funcs{
		UpdateFunc: func(e event.UpdateEvent) bool {
			// Ignore updates to CR status in which case metadata.Generation does not change.
			return e.ObjectOld.GetGeneration() != e.ObjectNew.GetGeneration()
		},
		DeleteFunc: func(e event.DeleteEvent) bool {
			// Evaluates to false if the object has been confirmed deleted.
			return !e.DeleteStateUnknown
		},
	}
}

// NewInstanaAgentReconciler initializes a new InstanaAgentReconciler instance
func NewInstanaAgentReconciler(
	client client.Client,
	scheme *runtime.Scheme,
	recorder record.EventRecorder,
	log logr.Logger,
) *InstanaAgentReconciler {
	return &InstanaAgentReconciler{
		client:       instanaclient.NewClient(client),
		recorder:     recorder,
		log:          log,
		chartRemover: helm.NewHelmReconciliation(scheme, log),
	}
}

type InstanaAgentReconciler struct {
	client       instanaclient.InstanaAgentClient
	recorder     record.EventRecorder
	log          logr.Logger
	chartRemover helm.DeprecatedInternalChartUninstaller
}

func (r *InstanaAgentReconciler) getAgent(ctx context.Context, req ctrl.Request) (
	*instanav1.InstanaAgent,
	reconcileReturn,
) {
	var agent instanav1.InstanaAgent

	switch err := r.client.Get(ctx, req.NamespacedName, &agent); {
	case k8serrors.IsNotFound(err):
		return nil, reconcileSuccess(ctrl.Result{})
	case !errors.Is(err, nil):
		return nil, reconcileFailure(err)
	default:
		return &agent, reconcileContinue()
	}
}

// TODO: Catch panics

// TODO: Update permissions here

// +kubebuilder:rbac:groups=agents.instana.io,resources=instanaagent,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apps,resources=daemonsets,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=pods;secrets;configmaps;services;serviceaccounts,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=agents.instana.io,resources=instanaagent/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=agents.instana.io,resources=instanaagent/finalizers,verbs=update
func (r *InstanaAgentReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	ctx = logr.NewContext(ctx, r.log)

	agent, res := r.getAgent(ctx, req)
	if res.suppliesReconcileResult() {
		return res.reconcileResult()
	}

}
