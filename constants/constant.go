package constants

import "time"

const LocalSource = "local"
const LocalFreeFormPath = "./configs/config.yaml"
const RemoteFreeformProfile = "config"
const Config = "config"
const CloudwatchNamespace = "AppConfig"
const CloudwatchErrorMetric = "ConfigurationError"
const CloudwatchErrorDimension = "Application"
const CloudwatchPutMetricInterval = 1 * time.Minute
const DeployedEnv = "env"
const DefaultPerPageLimit = 100
const Limit = "limit"
const Page = "page"
const CountRequired = "count_required"
