package daemonset

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	instanav1 "github.com/instana/instana-agent-operator/api/v1"
	"github.com/instana/instana-agent-operator/pkg/map_defaulter"
	"github.com/instana/instana-agent-operator/pkg/optional"
)

// TODO: Cleanup these and tests in other files

func TestDaemonSetBuilder_getPodTemplateLabels(t *testing.T) {
	t.Run(
		"agent_mode_unset", func(t *testing.T) {
			const name = "soijdfoijsfdoij"

			ctrl := gomock.NewController(t)

			expected := map[string]string{
				"adsf":      "eroinsvd",
				"osdgoiego": "rwuriunsv",
				"e8uriunv":  "rrudsiu",
			}

			podSelector := NewMockPodSelectorLabelGenerator(ctrl)
			podSelector.EXPECT().GetPodLabels(
				map[string]string{
					"instana/agent-mode": string(instanav1.APM),
				},
			).Return(expected)

			d := &daemonSetBuilder{
				InstanaAgent: &instanav1.InstanaAgent{
					ObjectMeta: metav1.ObjectMeta{
						Name: name,
					},
				},
				PodSelectorLabelGenerator: podSelector,
			}

			actual := d.getPodTemplateLabels()

			assertions := require.New(t)

			assertions.Equal(expected, actual)
		},
	)
	t.Run(
		"agent_mode_set_by_user", func(t *testing.T) {
			const name = "soijdfoijsfdoij"

			ctrl := gomock.NewController(t)

			expected := map[string]string{
				"adsf":      "eroinsvd",
				"osdgoiego": "rwuriunsv",
				"e8uriunv":  "rrudsiu",
			}

			podSelector := NewMockPodSelectorLabelGenerator(ctrl)
			podSelector.EXPECT().GetPodLabels(
				map[string]string{
					"instana/agent-mode": string(instanav1.KUBERNETES),
				},
			).Return(expected)

			d := &daemonSetBuilder{
				InstanaAgent: &instanav1.InstanaAgent{
					ObjectMeta: metav1.ObjectMeta{
						Name: name,
					},
					Spec: instanav1.InstanaAgentSpec{
						Agent: instanav1.BaseAgentSpec{
							Mode: instanav1.KUBERNETES,
						},
					},
				},
				PodSelectorLabelGenerator: podSelector,
			}

			actual := d.getPodTemplateLabels()

			assertions := require.New(t)

			assertions.Equal(expected, actual)
		},
	)
	t.Run(
		"agent_mode_unset_with_user_given_pod_labels", func(t *testing.T) {
			const name = "soijdfoijsfdoij"

			ctrl := gomock.NewController(t)

			expected := map[string]string{
				"adsf":      "eroinsvd",
				"osdgoiego": "rwuriunsv",
				"e8uriunv":  "rrudsiu",
			}

			podSelector := NewMockPodSelectorLabelGenerator(ctrl)
			podSelector.EXPECT().GetPodLabels(
				map[string]string{
					"asdfasdf":           "eoisdgoinv",
					"reoirionv":          "98458hgoisjdf",
					"instana/agent-mode": string(instanav1.APM),
				},
			).Return(expected)

			d := &daemonSetBuilder{
				InstanaAgent: &instanav1.InstanaAgent{
					ObjectMeta: metav1.ObjectMeta{
						Name: name,
					},
					Spec: instanav1.InstanaAgentSpec{
						Agent: instanav1.BaseAgentSpec{
							Pod: instanav1.AgentPodSpec{
								Labels: map[string]string{
									"asdfasdf":  "eoisdgoinv",
									"reoirionv": "98458hgoisjdf",
								},
							},
						},
					},
				},
				PodSelectorLabelGenerator: podSelector,
			}

			actual := d.getPodTemplateLabels()

			assertions := require.New(t)

			assertions.Equal(expected, actual)
		},
	)
	t.Run(
		"agent_mode_set_by_user_with_user_given_pod_labels", func(t *testing.T) {
			const name = "soijdfoijsfdoij"

			ctrl := gomock.NewController(t)

			expected := map[string]string{
				"adsf":      "eroinsvd",
				"osdgoiego": "rwuriunsv",
				"e8uriunv":  "rrudsiu",
			}

			podSelector := NewMockPodSelectorLabelGenerator(ctrl)
			podSelector.EXPECT().GetPodLabels(
				map[string]string{
					"asdfasdf":           "eoisdgoinv",
					"reoirionv":          "98458hgoisjdf",
					"instana/agent-mode": string(instanav1.KUBERNETES),
				},
			).Return(expected)

			d := &daemonSetBuilder{
				InstanaAgent: &instanav1.InstanaAgent{
					ObjectMeta: metav1.ObjectMeta{
						Name: name,
					},
					Spec: instanav1.InstanaAgentSpec{
						Agent: instanav1.BaseAgentSpec{
							Mode: instanav1.KUBERNETES,
							Pod: instanav1.AgentPodSpec{
								Labels: map[string]string{
									"asdfasdf":  "eoisdgoinv",
									"reoirionv": "98458hgoisjdf",
								},
							},
						},
					},
				},
				PodSelectorLabelGenerator: podSelector,
			}

			actual := d.getPodTemplateLabels()

			assertions := require.New(t)

			assertions.Equal(expected, actual)
		},
	)

}

func TestDaemonSetBuilder_getPodTemplateAnnotations(t *testing.T) {
	t.Run(
		"no_user_provided_annotations", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			assertions := require.New(t)

			agent := instanav1.InstanaAgent{
				Spec: instanav1.InstanaAgentSpec{
					Cluster: instanav1.Name{
						Name: "oawgeoieg",
					},
				},
			}

			const expectedHash = "49845soidghoijw09"

			hasher := NewMockJsonHasher(ctrl)
			hasher.EXPECT().HashJsonOrDie(gomock.Eq(&agent.Spec)).Return(expectedHash)

			db := &daemonSetBuilder{
				InstanaAgent: &agent,
				JsonHasher:   hasher,
			}

			actual := db.getPodTemplateAnnotations()
			assertions.Equal(
				map[string]string{
					"instana-configuration-hash": expectedHash,
				}, actual,
			)
		},
	)
	t.Run(
		"with_user_provided_annotations", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			assertions := require.New(t)

			agent := instanav1.InstanaAgent{
				Spec: instanav1.InstanaAgentSpec{
					Cluster: instanav1.Name{
						Name: "oawgeoieg",
					},
					Agent: instanav1.BaseAgentSpec{
						Pod: instanav1.AgentPodSpec{
							Annotations: map[string]string{
								"498hroihsg":             "4589fdoighjsoijs",
								"flkje489h309sd":         "oie409ojifg",
								"4509ufdoigjselkjweoihg": "g059pojw9jwpoijd",
							},
						},
					},
				},
			}

			const expectedHash = "49845soidghoijw09"

			hasher := NewMockJsonHasher(ctrl)
			hasher.EXPECT().HashJsonOrDie(gomock.Eq(&agent.Spec)).Return(expectedHash)

			db := &daemonSetBuilder{
				InstanaAgent: &agent,
				JsonHasher:   hasher,
			}

			actual := db.getPodTemplateAnnotations()
			assertions.Equal(
				map[string]string{
					"instana-configuration-hash": expectedHash,
					"498hroihsg":                 "4589fdoighjsoijs",
					"flkje489h309sd":             "oie409ojifg",
					"4509ufdoigjselkjweoihg":     "g059pojw9jwpoijd",
				}, actual,
			)
		},
	)
}

func TestDaemonSetBuilder_getImagePullSecrets(t *testing.T) {
	t.Run(
		"no_user_secrets_and_image_not_from_instana_io", func(t *testing.T) {
			assertions := require.New(t)

			db := &daemonSetBuilder{
				InstanaAgent: &instanav1.InstanaAgent{},
			}

			actual := db.getImagePullSecrets()
			assertions.Empty(actual)
		},
	)
	t.Run(
		"with_user_secrets_and_image_not_from_instana_io", func(t *testing.T) {
			assertions := require.New(t)

			db := &daemonSetBuilder{
				InstanaAgent: &instanav1.InstanaAgent{
					Spec: instanav1.InstanaAgentSpec{
						Agent: instanav1.BaseAgentSpec{
							ImageSpec: instanav1.ImageSpec{
								PullSecrets: []corev1.LocalObjectReference{
									{
										Name: "oirewigojsdf",
									},
									{
										Name: "o4gpoijsfd",
									},
									{
										Name: "po5hpojdfijs",
									},
								},
							},
						},
					},
				},
			}

			actual := db.getImagePullSecrets()

			assertions.Equal(
				[]corev1.LocalObjectReference{
					{
						Name: "oirewigojsdf",
					},
					{
						Name: "o4gpoijsfd",
					},
					{
						Name: "po5hpojdfijs",
					},
				},
				actual,
			)
		},
	)
	t.Run(
		"no_user_secrets_and_image_is_from_instana_io", func(t *testing.T) {
			assertions := require.New(t)

			db := &daemonSetBuilder{
				InstanaAgent: &instanav1.InstanaAgent{
					Spec: instanav1.InstanaAgentSpec{
						Agent: instanav1.BaseAgentSpec{
							ImageSpec: instanav1.ImageSpec{
								Name: "containers.instana.io/instana-agent",
							},
						},
					},
				},
			}

			actual := db.getImagePullSecrets()
			assertions.Equal(
				[]corev1.LocalObjectReference{
					{
						Name: "containers-instana-io",
					},
				},
				actual,
			)
		},
	)
	t.Run(
		"with_user_secrets_and_image_is_from_instana_io", func(t *testing.T) {
			assertions := require.New(t)

			db := &daemonSetBuilder{
				InstanaAgent: &instanav1.InstanaAgent{
					Spec: instanav1.InstanaAgentSpec{
						Agent: instanav1.BaseAgentSpec{
							ImageSpec: instanav1.ImageSpec{
								Name: "containers.instana.io/instana-agent",
								PullSecrets: []corev1.LocalObjectReference{
									{
										Name: "oirewigojsdf",
									},
									{
										Name: "o4gpoijsfd",
									},
									{
										Name: "po5hpojdfijs",
									},
								},
							},
						},
					},
				},
			}

			actual := db.getImagePullSecrets()

			assertions.Equal(
				[]corev1.LocalObjectReference{
					{
						Name: "oirewigojsdf",
					},
					{
						Name: "o4gpoijsfd",
					},
					{
						Name: "po5hpojdfijs",
					},
					{
						Name: "containers-instana-io",
					},
				},
				actual,
			)
		},
	)
}

func TestDaemonSetBuilder_getEnvVars(t *testing.T) {
	assertions := require.New(t)

	userProvidedEnv := map[string]string{
		"foo":   "bar",
		"hello": "world",
		"eodgh": "oijdsgnso",
	}

	db := NewDaemonSetBuilder(
		&instanav1.InstanaAgent{
			Spec: instanav1.InstanaAgentSpec{
				Agent: instanav1.BaseAgentSpec{
					Env: userProvidedEnv,
				},
			},
		},
		false,
	).(*daemonSetBuilder)
	res := db.getEnvVars()

	assertions.Len(res, 19+len(userProvidedEnv))
}

func TestDaemonSetBuilder_getResourceRequirements(t *testing.T) {
	metaAssertions := require.New(t)

	type testParams struct {
		providedMemRequest string
		providedCpuRequest string
		providedMemLimit   string
		providedCpuLimit   string

		expectedMemRequest string
		expectedCpuRequest string
		expectedMemLimit   string
		expectedCpuLimit   string
	}

	tests := make([]testParams, 0, 16)
	for _, providedMemRequest := range []string{"", "123Mi"} {
		for _, providedCpuRequest := range []string{"", "1.2"} {
			for _, providedMemLimit := range []string{"", "456Mi"} {
				for _, providedCpuLimit := range []string{"", "4.5"} {
					tests = append(
						tests, testParams{
							expectedMemRequest: optional.Of(providedMemRequest).GetOrDefault("512Mi"),
							expectedCpuRequest: optional.Of(providedCpuRequest).GetOrDefault("0.5"),
							expectedMemLimit:   optional.Of(providedMemLimit).GetOrDefault("768Mi"),
							expectedCpuLimit:   optional.Of(providedCpuLimit).GetOrDefault("1.5"),

							providedMemRequest: providedMemRequest,
							providedCpuRequest: providedCpuRequest,
							providedMemLimit:   providedMemLimit,
							providedCpuLimit:   providedCpuLimit,
						},
					)
				}
			}
		}
	}

	metaAssertions.Len(tests, 16)

	for _, test := range tests {
		t.Run(
			fmt.Sprintf("%+v", test), func(t *testing.T) {
				assertions := require.New(t)

				provided := corev1.ResourceRequirements{}

				setIfNotEmpty := func(providedVal string, key corev1.ResourceName, resourceList *corev1.ResourceList) {
					if providedVal != "" {
						map_defaulter.NewMapDefaulter((*map[corev1.ResourceName]resource.Quantity)(resourceList)).SetIfEmpty(
							key,
							resource.MustParse(providedVal),
						)
					}
				}

				setIfNotEmpty(test.providedMemLimit, corev1.ResourceMemory, &provided.Limits)
				setIfNotEmpty(test.providedCpuLimit, corev1.ResourceCPU, &provided.Limits)
				setIfNotEmpty(test.providedMemRequest, corev1.ResourceMemory, &provided.Requests)
				setIfNotEmpty(test.providedCpuRequest, corev1.ResourceCPU, &provided.Requests)

				db := &daemonSetBuilder{
					InstanaAgent: &instanav1.InstanaAgent{
						Spec: instanav1.InstanaAgentSpec{
							Agent: instanav1.BaseAgentSpec{
								Pod: instanav1.AgentPodSpec{
									ResourceRequirements: provided,
								},
							},
						},
					},
				}
				actual := db.getResourceRequirements()

				assertions.Equal(resource.MustParse(test.expectedMemLimit), actual.Limits[corev1.ResourceMemory])
				assertions.Equal(resource.MustParse(test.expectedCpuLimit), actual.Limits[corev1.ResourceCPU])
				assertions.Equal(resource.MustParse(test.expectedMemRequest), actual.Requests[corev1.ResourceMemory])
				assertions.Equal(resource.MustParse(test.expectedCpuRequest), actual.Requests[corev1.ResourceCPU])
			},
		)
	}
}
