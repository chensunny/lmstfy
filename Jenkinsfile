#!/usr/bin/env groovy
// Centralize Jenkins Pipelines configuration using Shared Libraries
// https://github.com/AfterShip/deployment-jenkins-pipeline-library-ci-aftership-org
@Library("jenkins-pipeline-library@automation") _
entry {
	isDryRun = false
	flow = "golang"
	configInfo = [
	      deploymentGroup          : "aftership",
<<<<<<< HEAD
          appName                  : "lmstfy.automizelyapi.org_main",
	      gitRepoName              : "lmstfy.automizelyapi.org_main.git",
	      chartName                : "golang-http",
	      essentialDockerImage     : "golang-essential",
	      essentialTag             : "golang-1.14.7",
	      requireStaticAsset       : false,
	      uploadAssetCredential    : "",
          domainType               : "automizelyapi.org",
	      unitTest                 : "",
	      integrationTest          : "",
	      useEnvironmentVariable   : true,
	      hasReleaseEnvironment    : false,
	      hasStagingEnvironment    : false,
	      hasProductionEnvironment : true,
	      useHelmReleaseShortName  : true
=======
	      appName                  : "lmstfy.automizelyapi.org_main",
	      gitRepoName              : "lmstfy.automizelyapi.org_main.git",
	      chartName                : "golang-http",
	      essentialDockerImage     : "golang-essential",
	      essentialTag             : "golang-1.13.5",
	      requireStaticAsset       : false,
	      uploadAssetCredential    : "",
	      domainType               : "automizelyapi.org",
	      unitTest                 : "test/unit_test.sh",
	      integrationTest          : "test/integration_test.sh",
	      useEnvironmentVariable   : true,
	      hasStagingEnvironment    : true,
	      hasProductionEnvironment : true,
	      useNPM                   : false,
		  useHelmReleaseShortName  : true
>>>>>>> update webservice Jenkinsfile and Dockerfile
	    ]
}
