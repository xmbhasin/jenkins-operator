package plugins

const (
	configurationAsCodePlugin           = "configuration-as-code:1903.v004d55388f30"
	gitPlugin                           = "git:5.7.0"
	jobDslPlugin                        = "job-dsl:1.89"
	kubernetesPlugin                    = "kubernetes:4295.v7fa_01b_309c95"
	kubernetesCredentialsProviderPlugin = "kubernetes-credentials-provider:1.262.v2670ef7ea_0c5"
	workflowAggregatorPlugin            = "workflow-aggregator:600.vb_57cdd26fdd7"
	workflowJobPlugin                   = "workflow-job:1472.ve4d5eca_143c4f"
)

// basePluginsList contains plugins to install by operator.
var basePluginsList = []Plugin{
	Must(New(configurationAsCodePlugin)),
	Must(New(gitPlugin)),
	Must(New(jobDslPlugin)),
	Must(New(kubernetesPlugin)),
	Must(New(kubernetesCredentialsProviderPlugin)),
	// Must(New(workflowJobPlugin)),
	Must(New(workflowAggregatorPlugin)),
}

// BasePlugins returns list of plugins to install by operator.
func BasePlugins() []Plugin {
	return basePluginsList
}
