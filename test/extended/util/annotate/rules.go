package main

// Rules defined here are additive to the rules already defined for
// kube e2e tests in openshift/kubernetes. The kube rules are
// vendored via the following file:
//
//   vendor/k8s.io/kubernetes/openshift-hack/e2e/annotate/rules.go
//
// Rules that are needed to pass the upstream e2e test suite in a
// "default OCP CI" configuration (eg, AWS or GCP, openshift-sdn) must
// be added to openshift/kubernetes to allow CI to pass there, and
// then vendored back into origin. Rules that only apply to
// "non-default" configurations (other clouds, other network
// providers) should be added here.

var (
	testMaps = map[string][]string{
		// tests that require a local host
		"[Local]": {
			// Doesn't work on scaled up clusters
			`\[Feature:ImagePrune\]`,
		},
		// alpha features that are not gated
		"[Disabled:Alpha]": {},
		// tests for features that are not implemented in openshift
		"[Disabled:Unimplemented]": {},
		// tests that rely on special configuration that we do not yet support
		"[Disabled:SpecialConfig]": {},
		// tests that rely on special configuration that we do not yet support
		// tests that are known broken and need to be fixed upstream or in openshift
		// always add an issue here
		"[Disabled:Broken]": {
			`should idle the service and DeploymentConfig properly`,       // idling with a single service and DeploymentConfig
			`should answer endpoint and wildcard queries for the cluster`, // currently not supported by dns operator https://github.com/openshift/cluster-dns-operator/issues/43

			// https://bugzilla.redhat.com/show_bug.cgi?id=1908645
			`\[sig-network\] Networking Granular Checks: Services should function for service endpoints using hostNetwork`,
			`\[sig-network\] Networking Granular Checks: Services should function for pod-Service\(hostNetwork\)`,

			// https://bugzilla.redhat.com/show_bug.cgi?id=1952460
			`\[sig-network\] Firewall rule control plane should not expose well-known ports`,

			// https://bugzilla.redhat.com/show_bug.cgi?id=1952457
			`\[sig-node\] crictl should be able to run crictl on the node`,

			// https://bugzilla.redhat.com/show_bug.cgi?id=1945091
			`\[Feature:GenericEphemeralVolume\]`,

			// https://bugzilla.redhat.com/show_bug.cgi?id=1953478
			`\[sig-storage\] Dynamic Provisioning Invalid AWS KMS key should report an error and create no PV`,

			// https://bugzilla.redhat.com/show_bug.cgi?id=1957894
			`\[sig-node\] Container Runtime blackbox test when running a container with a new image should be able to pull from private registry with secret`,

			// tests disabled due to API changes coming from k8s 1.22:
			// - coming from ServiceLBNodePortControl changes https://github.com/kubernetes/kubernetes/pull/100412
			// - https://bugzilla.redhat.com/show_bug.cgi?id=1986001 - sig-api-machinery
			`\[sig-api-machinery\] ResourceQuota should create a ResourceQuota and capture the life of a service`,
			`\[sig-api-machinery\] API data in etcd`,

			// https://bugzilla.redhat.com/show_bug.cgi?id=1989417
			`\[sig-cli\] oc adm storage-admin`,

			// https://bugzilla.redhat.com/show_bug.cgi?id=1989423
			`\[sig-network-edge\]\[Conformance\]\[Area:Networking\]\[Feature:Router\] The HAProxy router should be able to connect to a service that is idled because a GET on the route will unidle it`,

			// due to CSI migration the following tests needs to be disabled:
			// - https://bugzilla.redhat.com/show_bug.cgi?id=1985998 - sig-storage
			`\[sig-storage\] CSI Volumes \[Driver: csi-hostpath\] \[Testpattern: Dynamic PV \(block volmode\)\(allowExpansion\)\] volume-expand Verify if offline PVC expansion works`,
			`\[sig-storage\] CSI Volumes \[Driver: csi-hostpath\] \[Testpattern: Dynamic PV \(block volmode\)\(allowExpansion\)\] volume-expand should resize volume when PVC is edited while pod is using it`,
			`\[sig-storage\] CSI Volumes \[Driver: csi-hostpath\] \[Testpattern: Dynamic PV \(block volmode\)\] provisioning should provision storage with pvc data source`,
			`\[sig-storage\] CSI Volumes \[Driver: csi-hostpath\] \[Testpattern: Dynamic PV \(block volmode\)\] provisioning should provision storage with snapshot data source`,
			`\[sig-storage\] CSI Volumes \[Driver: csi-hostpath\] \[Testpattern: Dynamic PV \(block volmode\)\] volume-expand should not allow expansion of pvcs without AllowVolumeExpansion property`,
			`\[sig-storage\] CSI Volumes \[Driver: csi-hostpath\] \[Testpattern: Dynamic PV \(block volmode\)\] volumeMode should not mount / map unused volumes in a pod`,
			`\[sig-storage\] CSI Volumes \[Driver: csi-hostpath\] \[Testpattern: Dynamic PV \(block volmode\)\] volumes should store data`,
			`\[sig-storage\] CSI Volumes \[Driver: csi-hostpath\] \[Testpattern: Dynamic PV \(default fs\)\(allowExpansion\)\] volume-expand Verify if offline PVC expansion works`,
			`\[sig-storage\] CSI Volumes \[Driver: csi-hostpath\] \[Testpattern: Dynamic PV \(default fs\)\(allowExpansion\)\] volume-expand should resize volume when PVC is edited while pod is using it`,
			`\[sig-storage\] CSI Volumes \[Driver: csi-hostpath\] \[Testpattern: Dynamic PV \(default fs\)\] provisioning should provision storage with pvc data source`,
			`\[sig-storage\] CSI Volumes \[Driver: csi-hostpath\] \[Testpattern: Dynamic PV \(default fs\)\] provisioning should provision storage with snapshot data source`,
			`\[sig-storage\] CSI Volumes \[Driver: csi-hostpath\] \[Testpattern: Dynamic PV \(default fs\)\] subPath should be able to unmount after the subpath directory is deleted`,
			`\[sig-storage\] CSI Volumes \[Driver: csi-hostpath\] \[Testpattern: Dynamic PV \(default fs\)\] subPath should support existing directories when readOnly specified in the volumeSource`,
			`\[sig-storage\] CSI Volumes \[Driver: csi-hostpath\] \[Testpattern: Dynamic PV \(default fs\)\] subPath should support existing directory`,
			`\[sig-storage\] CSI Volumes \[Driver: csi-hostpath\] \[Testpattern: Dynamic PV \(default fs\)\] subPath should support existing single file`,
			`\[sig-storage\] CSI Volumes \[Driver: csi-hostpath\] \[Testpattern: Dynamic PV \(default fs\)\] subPath should support file as subpath`,
			`\[sig-storage\] CSI Volumes \[Driver: csi-hostpath\] \[Testpattern: Dynamic PV \(default fs\)\] subPath should support non-existent path`,
			`\[sig-storage\] CSI Volumes \[Driver: csi-hostpath\] \[Testpattern: Dynamic PV \(default fs\)\] subPath should support readOnly directory specified in the volumeMount`,
			`\[sig-storage\] CSI Volumes \[Driver: csi-hostpath\] \[Testpattern: Dynamic PV \(default fs\)\] subPath should support readOnly file specified in the volumeMount`,
			`\[sig-storage\] CSI Volumes \[Driver: csi-hostpath\] \[Testpattern: Dynamic PV \(default fs\)\] volume-expand should not allow expansion of pvcs without AllowVolumeExpansion property`,
			`\[sig-storage\] CSI Volumes \[Driver: csi-hostpath\] \[Testpattern: Dynamic PV \(default fs\)\] volumes should store data`,
			`\[sig-storage\] CSI Volumes \[Driver: csi-hostpath\] \[Testpattern: Dynamic PV \(filesystem volmode\)\] volumeLimits should support volume limits`,
			`\[sig-storage\] CSI Volumes \[Driver: csi-hostpath\] \[Testpattern: Dynamic PV \(filesystem volmode\)\] volumeMode should not mount / map unused volumes in a pod`,
			`\[sig-storage\] CSI Volumes \[Driver: csi-hostpath\] \[Testpattern: Dynamic Snapshot \(delete policy\)\] snapshottable\[Feature:VolumeSnapshotDataSource\] volume snapshot controller  should check snapshot fields, check restore correctly works after modifying source data, check deletion`,
			`\[sig-storage\] CSI Volumes \[Driver: csi-hostpath\] \[Testpattern: Dynamic Snapshot \(retain policy\)\] snapshottable\[Feature:VolumeSnapshotDataSource\] volume snapshot controller  should check snapshot fields, check restore correctly works after modifying source data, check deletion`,
			`\[sig-storage\] CSI Volumes \[Driver: csi-hostpath\] \[Testpattern: Generic Ephemeral-volume \(default fs\) \(immediate-binding\)\] ephemeral should create read-only inline ephemeral volume`,
			`\[sig-storage\] CSI Volumes \[Driver: csi-hostpath\] \[Testpattern: Generic Ephemeral-volume \(default fs\) \(immediate-binding\)\] ephemeral should create read/write inline ephemeral volume`,
			`\[sig-storage\] CSI Volumes \[Driver: csi-hostpath\] \[Testpattern: Generic Ephemeral-volume \(default fs\) \(immediate-binding\)\] ephemeral should support multiple inline ephemeral volumes`,
			`\[sig-storage\] CSI Volumes \[Driver: csi-hostpath\] \[Testpattern: Generic Ephemeral-volume \(default fs\) \(immediate-binding\)\] ephemeral should support two pods which share the same volume`,
			`\[sig-storage\] CSI Volumes \[Driver: csi-hostpath\] \[Testpattern: Generic Ephemeral-volume \(default fs\) \(late-binding\)\] ephemeral should create read-only inline ephemeral volume`,
			`\[sig-storage\] CSI Volumes \[Driver: csi-hostpath\] \[Testpattern: Generic Ephemeral-volume \(default fs\) \(late-binding\)\] ephemeral should create read/write inline ephemeral volume`,
			`\[sig-storage\] CSI Volumes \[Driver: csi-hostpath\] \[Testpattern: Generic Ephemeral-volume \(default fs\) \(late-binding\)\] ephemeral should support multiple inline ephemeral volumes`,
			`\[sig-storage\] CSI Volumes \[Driver: csi-hostpath\] \[Testpattern: Generic Ephemeral-volume \(default fs\) \(late-binding\)\] ephemeral should support two pods which share the same volume`,
			`\[sig-storage\] CSI Volumes \[Driver: csi-hostpath\] \[Testpattern: Pre-provisioned Snapshot \(delete policy\)\] snapshottable\[Feature:VolumeSnapshotDataSource\] volume snapshot controller  should check snapshot fields, check restore correctly works after modifying source data, check deletion`,
			`\[sig-storage\] CSI Volumes \[Driver: csi-hostpath\] \[Testpattern: Pre-provisioned Snapshot \(retain policy\)\] snapshottable\[Feature:VolumeSnapshotDataSource\] volume snapshot controller  should check snapshot fields, check restore correctly works after modifying source data, check deletion`,
			`\[sig-storage\] CSI mock volume CSI Volume expansion should expand volume by restarting pod if attach=on, nodeExpansion=on`,
			`\[sig-storage\] CSI mock volume CSI Volume expansion should expand volume without restarting pod if nodeExpansion=off`,
			`\[sig-storage\] CSI mock volume CSI Volume expansion should not expand volume if resizingOnDriver=off, resizingOnSC=on`,
			`\[sig-storage\] CSI mock volume CSI attach test using mock driver should preserve attachment policy when no CSIDriver present`,
			`\[sig-storage\] CSI mock volume CSI attach test using mock driver should require VolumeAttach for drivers with attachment`,
			`\[sig-storage\] CSI mock volume CSI online volume expansion should expand volume without restarting pod if attach=on, nodeExpansion=on`,
			`\[sig-storage\] CSI mock volume CSI workload information using mock driver should be passed when podInfoOnMount=true`,
			`\[sig-storage\] CSI mock volume CSI workload information using mock driver should not be passed when CSIDriver does not exist`,
			`\[sig-storage\] CSI mock volume CSI workload information using mock driver should not be passed when podInfoOnMount=false`,
			`\[sig-storage\] CSI mock volume CSI workload information using mock driver should not be passed when podInfoOnMount=nil`,
			`\[sig-storage\] CSI mock volume CSIServiceAccountToken token should be plumbed down when csiServiceAccountTokenEnabled=true`,
			`\[sig-storage\] CSI mock volume CSIServiceAccountToken token should not be plumbed down when CSIDriver is not deployed`,
			`\[sig-storage\] CSI mock volume CSIServiceAccountToken token should not be plumbed down when csiServiceAccountTokenEnabled=false`,
			`\[sig-storage\] CSI mock volume CSIStorageCapacity CSIStorageCapacity disabled`,
			`\[sig-storage\] CSI mock volume CSIStorageCapacity CSIStorageCapacity unused`,
			`\[sig-storage\] CSI mock volume CSIStorageCapacity CSIStorageCapacity used, have capacity`,

			// The new NetworkPolicy test suite is extremely resource
			// intensive and causes itself and other concurrently-running
			// tests to be flaky.
			// https://bugzilla.redhat.com/show_bug.cgi?id=1980141
			`\[sig-network\] Netpol `,
		},
		// tests that may work, but we don't support them
		"[Disabled:Unsupported]": {},
		// tests too slow to be part of conformance
		"[Slow]": {},
		// tests that are known flaky
		"[Flaky]": {
			`openshift mongodb replication creating from a template`, // flaking on deployment

			// BZ https://bugzilla.redhat.com/show_bug.cgi?id=1991507
			`Kubectl client Simple pod should return command exit codes`,
		},
		// tests that must be run without competition
		"[Serial]":        {},
		"[Skipped:azure]": {},
		"[Skipped:ovirt]": {},
		"[Skipped:gce]":   {},

		"[Skipped:SingleReplicaTopology]": {
			`\[sig-apps\] Daemon set \[Serial\] should rollback without unnecessary restarts \[Conformance\]`,
			`\[sig-node\] NoExecuteTaintManager Single Pod \[Serial\] doesn't evict pod with tolerations from tainted nodes`,
			`\[sig-node\] NoExecuteTaintManager Single Pod \[Serial\] eventually evict pod with finite tolerations from tainted nodes`,
			`\[sig-node\] NoExecuteTaintManager Single Pod \[Serial\] evicts pods from tainted nodes`,
			`\[sig-node\] NoExecuteTaintManager Single Pod \[Serial\] removing taint cancels eviction \[Disruptive\] \[Conformance\]`,
			`\[sig-node\] NoExecuteTaintManager Multiple Pods \[Serial\] evicts pods with minTolerationSeconds \[Disruptive\] \[Conformance\]`,
			`\[sig-node\] NoExecuteTaintManager Multiple Pods \[Serial\] only evicts pods without tolerations from tainted nodes`,
			`\[sig-cli\] Kubectl client Kubectl taint \[Serial\] should remove all the taints with the same key off a node`,
		},

		// Tests that don't pass on disconnected, either due to requiring
		// internet access for GitHub (e.g. many of the s2i builds), or
		// because of pullthrough not supporting ICSP (https://bugzilla.redhat.com/show_bug.cgi?id=1918376)
		"[Skipped:Disconnected]": {
			// Internet access required
			`\[sig-builds\]\[Feature:Builds\] clone repository using git:// protocol  should clone using git:// if no proxy is configured`,
			`\[sig-builds\]\[Feature:Builds\] result image should have proper labels set  S2I build from a template should create a image from "test-s2i-build.json" template with proper Docker labels`,
			`\[sig-builds\]\[Feature:Builds\] s2i build with a quota  Building from a template should create an s2i build with a quota and run it`,
			`\[sig-builds\]\[Feature:Builds\] s2i build with a root user image should create a root build and pass with a privileged SCC`,
			`\[sig-builds\]\[Feature:Builds\]\[timing\] capture build stages and durations  should record build stages and durations for docker`,
			`\[sig-builds\]\[Feature:Builds\]\[timing\] capture build stages and durations  should record build stages and durations for s2i`,
			`\[sig-builds\]\[Feature:Builds\]\[valueFrom\] process valueFrom in build strategy environment variables  should successfully resolve valueFrom in s2i build environment variables`,
			`\[sig-builds\]\[Feature:Builds\]\[volumes\] should mount given secrets and configmaps into the build pod for source strategy builds`,
			`\[sig-builds\]\[Feature:Builds\]\[volumes\] should mount given secrets and configmaps into the build pod for docker strategy builds`,
			`\[sig-cli\] oc debug ensure it works with image streams`,
			`\[sig-cli\] oc builds complex build start-build`,
			`\[sig-cli\] oc builds complex build webhooks CRUD`,
			`\[sig-cli\] oc builds new-build`,
			`\[sig-devex\] check registry.redhat.io is available and samples operator can import sample imagestreams run sample related validations`,
			`\[sig-devex\]\[Feature:Templates\] templateinstance readiness test  should report failed soon after an annotated objects has failed`,
			`\[sig-devex\]\[Feature:Templates\] templateinstance readiness test  should report ready soon after all annotated objects are ready`,
			`\[sig-operator\] an end user can use OLM can subscribe to the operator`,
			`\[sig-network\] Networking should provide Internet connection for containers`,

			// Need to access non-cached images like ruby and mongodb
			`\[sig-apps\]\[Feature:DeploymentConfig\] deploymentconfigs with multiple image change triggers should run a successful deployment with a trigger used by different containers`,
			`\[sig-apps\]\[Feature:DeploymentConfig\] deploymentconfigs with multiple image change triggers should run a successful deployment with multiple triggers`,

			// ICSP
			`\[sig-apps\]\[Feature:DeploymentConfig\] deploymentconfigs  should adhere to Three Laws of Controllers`,
			`\[sig-apps\]\[Feature:DeploymentConfig\] deploymentconfigs adoption will orphan all RCs and adopt them back when recreated`,
			`\[sig-apps\]\[Feature:DeploymentConfig\] deploymentconfigs generation should deploy based on a status version bump`,
			`\[sig-apps\]\[Feature:DeploymentConfig\] deploymentconfigs keep the deployer pod invariant valid should deal with cancellation after deployer pod succeeded`,
			`\[sig-apps\]\[Feature:DeploymentConfig\] deploymentconfigs paused should disable actions on deployments`,
			`\[sig-apps\]\[Feature:DeploymentConfig\] deploymentconfigs rolled back should rollback to an older deployment`,
			`\[sig-apps\]\[Feature:DeploymentConfig\] deploymentconfigs should respect image stream tag reference policy resolve the image pull spec`,
			`\[sig-apps\]\[Feature:DeploymentConfig\] deploymentconfigs viewing rollout history should print the rollout history`,
			`\[sig-apps\]\[Feature:DeploymentConfig\] deploymentconfigs when changing image change trigger should successfully trigger from an updated image`,
			`\[sig-apps\]\[Feature:DeploymentConfig\] deploymentconfigs when run iteratively should only deploy the last deployment`,
			`\[sig-apps\]\[Feature:DeploymentConfig\] deploymentconfigs when tagging images should successfully tag the deployed image`,
			`\[sig-apps\]\[Feature:DeploymentConfig\] deploymentconfigs with custom deployments should run the custom deployment steps`,
			`\[sig-apps\]\[Feature:DeploymentConfig\] deploymentconfigs with enhanced status should include various info in status`,
			`\[sig-apps\]\[Feature:DeploymentConfig\] deploymentconfigs with env in params referencing the configmap should expand the config map key to a value`,
			`\[sig-apps\]\[Feature:DeploymentConfig\] deploymentconfigs with failing hook should get all logs from retried hooks`,
			`\[sig-apps\]\[Feature:DeploymentConfig\] deploymentconfigs with minimum ready seconds set should not transition the deployment to Complete before satisfied`,
			`\[sig-apps\]\[Feature:DeploymentConfig\] deploymentconfigs with revision history limits should never persist more old deployments than acceptable after being observed by the controller`,
			`\[sig-apps\]\[Feature:DeploymentConfig\] deploymentconfigs with test deployments should run a deployment to completion and then scale to zero`,
			`\[sig-apps\]\[Feature:DeploymentConfig\] deploymentconfigs won't deploy RC with unresolved images when patched with empty image`,
			`\[sig-apps\]\[Feature:Jobs\] Users should be able to create and run a job in a user project`,
			`\[sig-arch\] Managed cluster should should expose cluster services outside the cluster`,
			`\[sig-arch\]\[Early\] Managed cluster should start all core operators`,
			`\[sig-auth\]\[Feature:SecurityContextConstraints\]  TestPodDefaultCapabilities`,
			`\[sig-builds\]\[Feature:Builds\] Multi-stage image builds should succeed`,
			`\[sig-builds\]\[Feature:Builds\] Optimized image builds  should succeed`,
			`\[sig-builds\]\[Feature:Builds\] build can reference a cluster service  with a build being created from new-build should be able to run a build that references a cluster service`,
			`\[sig-builds\]\[Feature:Builds\] build have source revision metadata  started build should contain source revision information`,
			`\[sig-builds\]\[Feature:Builds\] build with empty source  started build should build even with an empty source in build config`,
			`\[sig-builds\]\[Feature:Builds\] build without output image  building from templates should create an image from a S2i template without an output image reference defined`,
			`\[sig-builds\]\[Feature:Builds\] build without output image  building from templates should create an image from a docker template without an output image reference defined`,
			`\[sig-builds\]\[Feature:Builds\] custom build with buildah  being created from new-build should complete build with custom builder image`,
			`\[sig-builds\]\[Feature:Builds\] imagechangetriggers  imagechangetriggers should trigger builds of all types`,
			`\[sig-builds\]\[Feature:Builds\] oc new-app  should fail with a --name longer than 58 characters`,
			`\[sig-builds\]\[Feature:Builds\] oc new-app  should succeed with a --name of 58 characters`,
			`\[sig-builds\]\[Feature:Builds\] oc new-app  should succeed with an imagestream`,
			`\[sig-builds\]\[Feature:Builds\] prune builds based on settings in the buildconfig  buildconfigs should have a default history limit set when created via the group api`,
			`\[sig-builds\]\[Feature:Builds\] prune builds based on settings in the buildconfig  should prune builds after a buildConfig change`,
			`\[sig-builds\]\[Feature:Builds\] prune builds based on settings in the buildconfig  should prune canceled builds based on the failedBuildsHistoryLimit setting`,
			`\[sig-builds\]\[Feature:Builds\] prune builds based on settings in the buildconfig  should prune completed builds based on the successfulBuildsHistoryLimit setting`,
			`\[sig-builds\]\[Feature:Builds\] prune builds based on settings in the buildconfig  should prune errored builds based on the failedBuildsHistoryLimit setting`,
			`\[sig-builds\]\[Feature:Builds\] prune builds based on settings in the buildconfig  should prune failed builds based on the failedBuildsHistoryLimit setting`,
			`\[sig-builds\]\[Feature:Builds\] result image should have proper labels set  Docker build from a template should create a image from "test-docker-build.json" template with proper Docker labels`,
			`\[sig-builds\]\[Feature:Builds\] verify /run filesystem contents  are writeable using a simple Docker Strategy Build`,
			`\[sig-builds\]\[Feature:Builds\] verify /run filesystem contents  do not have unexpected content using a simple Docker Strategy Build`,
			`\[sig-builds\]\[Feature:Builds\]\[pullsecret\] docker build using a pull secret  Building from a template should create a docker build that pulls using a secret run it`,
			`\[sig-builds\]\[Feature:Builds\]\[valueFrom\] process valueFrom in build strategy environment variables  should fail resolving unresolvable valueFrom in docker build environment variable references`,
			`\[sig-builds\]\[Feature:Builds\]\[valueFrom\] process valueFrom in build strategy environment variables  should fail resolving unresolvable valueFrom in sti build environment variable references`,
			`\[sig-builds\]\[Feature:Builds\]\[valueFrom\] process valueFrom in build strategy environment variables  should successfully resolve valueFrom in docker build environment variables`,
			`\[sig-cli\] CLI can run inside of a busybox container`,
			`\[sig-cli\] oc debug deployment configs from a build`,
			`\[sig-cli\] oc rsh specific flags should work well when access to a remote shell`,
			`\[sig-cli\] oc builds get buildconfig`,
			`\[sig-cli\] oc builds patch buildconfig`,
			`\[sig-cluster-lifecycle\] Pods cannot access the /config/master API endpoint`,
			`\[sig-imageregistry\]\[Feature:ImageAppend\] Image append should create images by appending them`,
			`\[sig-imageregistry\]\[Feature:ImageExtract\] Image extract should extract content from an image`,
			`\[sig-imageregistry\]\[Feature:ImageInfo\] Image info should display information about images`,
			`\[sig-imageregistry\]\[Feature:ImageLayers\] Image layer subresource should return layers from tagged images`,
			`\[sig-imageregistry\]\[Feature:ImageTriggers\] Annotation trigger reconciles after the image is overwritten`,
			`\[sig-imageregistry\]\[Feature:Image\] oc tag should change image reference for internal images`,
			`\[sig-imageregistry\]\[Feature:Image\] oc tag should work when only imagestreams api is available`,
			`\[sig-instrumentation\] Prometheus when installed on the cluster should have a AlertmanagerReceiversNotConfigured alert in firing state`,
			`\[sig-instrumentation\] Prometheus when installed on the cluster should have important platform topology metrics`,
			`\[sig-instrumentation\] Prometheus when installed on the cluster should have non-Pod host cAdvisor metrics`,
			`\[sig-instrumentation\] Prometheus when installed on the cluster should provide ingress metrics`,
			`\[sig-instrumentation\] Prometheus when installed on the cluster should provide named network metrics`,
			`\[sig-instrumentation\] Prometheus when installed on the cluster should report telemetry if a cloud.openshift.com token is present \[Late\]`,
			`\[sig-instrumentation\] Prometheus when installed on the cluster should start and expose a secured proxy and unsecured metrics`,
			`\[sig-instrumentation\] Prometheus when installed on the cluster shouldn't have failing rules evaluation`,
			`\[sig-instrumentation\] Prometheus when installed on the cluster shouldn't report any alerts in firing state apart from Watchdog and AlertmanagerReceiversNotConfigured \[Early\]`,
			`\[sig-instrumentation\] Prometheus when installed on the cluster when using openshift-sdn should be able to get the sdn ovs flows`,
			`\[sig-instrumentation\]\[Late\] Alerts should have a Watchdog alert in firing state the entire cluster run`,
			`\[sig-instrumentation\]\[Late\] Alerts shouldn't exceed the 500 series limit of total series sent via telemetry from each cluster`,
			`\[sig-instrumentation\]\[Late\] Alerts shouldn't report any alerts in firing or pending state apart from Watchdog and AlertmanagerReceiversNotConfigured and have no gaps in Watchdog firing`,
			`\[sig-instrumentation\]\[sig-builds\]\[Feature:Builds\] Prometheus when installed on the cluster should start and expose a secured proxy and verify build metrics`,
			`\[sig-network-edge\]\[Conformance\]\[Area:Networking\]\[Feature:Router\] The HAProxy router should be able to connect to a service that is idled because a GET on the route will unidle it`,
			`\[sig-network\]\[Feature:Router\] The HAProxy router should enable openshift-monitoring to pull metrics`,
			`\[sig-network\]\[Feature:Router\] The HAProxy router should expose a health check on the metrics port`,
			`\[sig-network\]\[Feature:Router\] The HAProxy router should expose prometheus metrics for a route`,
			`\[sig-network\]\[Feature:Router\] The HAProxy router should expose the profiling endpoints`,
			`\[sig-network\]\[Feature:Router\] The HAProxy router should override the route host for overridden domains with a custom value`,
			`\[sig-network\]\[Feature:Router\] The HAProxy router should override the route host with a custom value`,
			`\[sig-network\]\[Feature:Router\] The HAProxy router should respond with 503 to unrecognized hosts`,
			`\[sig-network\]\[Feature:Router\] The HAProxy router should run even if it has no access to update status`,
			`\[sig-network\]\[Feature:Router\] The HAProxy router should serve a route that points to two services and respect weights`,
			`\[sig-network\]\[Feature:Router\] The HAProxy router should serve routes that were created from an ingress`,
			`\[sig-network\]\[Feature:Router\] The HAProxy router should serve the correct routes when scoped to a single namespace and label set`,
			`\[sig-network\]\[Feature:Router\] The HAProxy router should set Forwarded headers appropriately`,
			`\[sig-network\]\[Feature:Router\] The HAProxy router should support reencrypt to services backed by a serving certificate automatically`,
			`\[sig-network\] Networking should provide Internet connection for containers \[Feature:Networking-IPv6\]`,
			`\[sig-node\] Managed cluster should report ready nodes the entire duration of the test run`,
			`\[sig-storage\]\[Late\] Metrics should report short attach times`,
			`\[sig-storage\]\[Late\] Metrics should report short mount times`,
		},

		// tests that don't pass under openshift-sdn NetworkPolicy mode are specified
		// in the rules file in openshift/kubernetes, not here.

		// tests that don't pass under openshift-sdn multitenant mode
		"[Skipped:Network/OpenShiftSDN/Multitenant]": {
			`\[Feature:NetworkPolicy\]`, // not compatible with multitenant mode
		},
		// tests that don't pass under OVN Kubernetes
		"[Skipped:Network/OVNKubernetes]": {
			// https://jira.coreos.com/browse/SDN-510: OVN-K doesn't support session affinity
			`\[sig-network\] Networking Granular Checks: Services should function for client IP based session affinity: http`,
			`\[sig-network\] Networking Granular Checks: Services should function for client IP based session affinity: udp`,
			`\[sig-network\] Services should be able to switch session affinity for NodePort service`,
			`\[sig-network\] Services should be able to switch session affinity for service with type clusterIP`,
			`\[sig-network\] Services should have session affinity work for NodePort service`,
			`\[sig-network\] Services should have session affinity work for service with type clusterIP`,
			`\[sig-network\] Services should have session affinity timeout work for NodePort service`,
			`\[sig-network\] Services should have session affinity timeout work for service with type clusterIP`,

			// https://bugzilla.redhat.com/show_bug.cgi?id=1972684 - [Feature:IPv6DualStack] tests are failing in dualstack
			// https://jira.coreos.com/browse/SDN-510: OVN-K doesn't support session affinity
			`\[sig-network\] \[Feature:IPv6DualStack\] \[LinuxOnly\] Granular Checks: Services Secondary IP Family should function for client IP based session affinity: http`,
			`\[sig-network\] \[Feature:IPv6DualStack\] \[LinuxOnly\] Granular Checks: Services Secondary IP Family should function for client IP based session affinity: udp`,
			`\[sig-network\] \[Feature:IPv6DualStack\] \[LinuxOnly\] should have ipv4 and ipv6 node podCIDRs`,

			// ovn-kubernetes does not support named ports
			`NetworkPolicy.*named port`,
		},
		"[Skipped:ibmcloud]": {
			// skip Gluster tests (not supported on ROKS worker nodes)
			// https://bugzilla.redhat.com/show_bug.cgi?id=1825009 - e2e: skip Glusterfs-related tests upstream for rhel7 worker nodes
			`\[Driver: gluster\]`,
			`GlusterFS`,
			`GlusterDynamicProvisioner`,

			// Nodes in ROKS have access to secrets in the cluster to handle encryption
			// https://bugzilla.redhat.com/show_bug.cgi?id=1825013 - ROKS: worker nodes have access to secrets in the cluster
			`\[sig-auth\] \[Feature:NodeAuthorizer\] Getting a non-existent configmap should exit with the Forbidden error, not a NotFound error`,
			`\[sig-auth\] \[Feature:NodeAuthorizer\] Getting a non-existent secret should exit with the Forbidden error, not a NotFound error`,
			`\[sig-auth\] \[Feature:NodeAuthorizer\] Getting a secret for a workload the node has access to should succeed`,
			`\[sig-auth\] \[Feature:NodeAuthorizer\] Getting an existing configmap should exit with the Forbidden error`,
			`\[sig-auth\] \[Feature:NodeAuthorizer\] Getting an existing secret should exit with the Forbidden error`,

			// Access to node external address is blocked from pods within a ROKS cluster by Calico
			// https://bugzilla.redhat.com/show_bug.cgi?id=1825016 - e2e: NodeAuthenticator tests use both external and internal addresses for node
			`\[sig-auth\] \[Feature:NodeAuthenticator\] The kubelet's main port 10250 should reject requests with no credentials`,
			`\[sig-auth\] \[Feature:NodeAuthenticator\] The kubelet can delegate ServiceAccount tokens to the API server`,

			// Calico is allowing the request to timeout instead of returning 'REFUSED'
			// https://bugzilla.redhat.com/show_bug.cgi?id=1825021 - ROKS: calico SDN results in a request timeout when accessing services with no endpoints
			`\[sig-network\] Services should be rejected when no endpoints exist`,

			// Mode returned by RHEL7 worker contains an extra character not expected by the test: dgtrwx vs dtrwx
			// https://bugzilla.redhat.com/show_bug.cgi?id=1825024 - e2e: Failing test - HostPath should give a volume the correct mode
			`\[sig-storage\] HostPath should give a volume the correct mode`,

			// Currently ibm-master-proxy-static and imbcloud-block-storage-plugin tolerate all taints
			// https://bugzilla.redhat.com/show_bug.cgi?id=1825027
			`\[Feature:Platform\] Managed cluster should ensure control plane operators do not make themselves unevictable`,
		},
	}

	// labelExcludes temporarily block tests out of a specific suite
	labelExcludes = map[string][]string{}
)
