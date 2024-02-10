package batch

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// AccessScope enumerates the values for access scope.
type AccessScope string

const (
	// Job Grants access to perform all operations on the job containing the task.
	Job AccessScope = "job"
)

// PossibleAccessScopeValues returns an array of possible values for the AccessScope const type.
func PossibleAccessScopeValues() []AccessScope {
	return []AccessScope{Job}
}

// AllocationState enumerates the values for allocation state.
type AllocationState string

const (
	// Resizing The pool is resizing; that is, compute nodes are being added to or removed from the pool.
	Resizing AllocationState = "resizing"
	// Steady The pool is not resizing. There are no changes to the number of nodes in the pool in progress. A
	// pool enters this state when it is created and when no operations are being performed on the pool to
	// change the number of nodes.
	Steady AllocationState = "steady"
	// Stopping The pool was resizing, but the user has requested that the resize be stopped, but the stop
	// request has not yet been completed.
	Stopping AllocationState = "stopping"
)

// PossibleAllocationStateValues returns an array of possible values for the AllocationState const type.
func PossibleAllocationStateValues() []AllocationState {
	return []AllocationState{Resizing, Steady, Stopping}
}

// AutoUserScope enumerates the values for auto user scope.
type AutoUserScope string

const (
	// Pool Specifies that the task runs as the common auto user account which is created on every node in a
	// pool.
	Pool AutoUserScope = "pool"
	// Task Specifies that the service should create a new user for the task.
	Task AutoUserScope = "task"
)

// PossibleAutoUserScopeValues returns an array of possible values for the AutoUserScope const type.
func PossibleAutoUserScopeValues() []AutoUserScope {
	return []AutoUserScope{Pool, Task}
}

// CachingType enumerates the values for caching type.
type CachingType string

const (
	// None The caching mode for the disk is not enabled.
	None CachingType = "none"
	// ReadOnly The caching mode for the disk is read only.
	ReadOnly CachingType = "readonly"
	// ReadWrite The caching mode for the disk is read and write.
	ReadWrite CachingType = "readwrite"
)

// PossibleCachingTypeValues returns an array of possible values for the CachingType const type.
func PossibleCachingTypeValues() []CachingType {
	return []CachingType{None, ReadOnly, ReadWrite}
}

// CertificateFormat enumerates the values for certificate format.
type CertificateFormat string

const (
	// Cer The certificate is a base64-encoded X.509 certificate.
	Cer CertificateFormat = "cer"
	// Pfx The certificate is a PFX (PKCS#12) formatted certificate or certificate chain.
	Pfx CertificateFormat = "pfx"
)

// PossibleCertificateFormatValues returns an array of possible values for the CertificateFormat const type.
func PossibleCertificateFormatValues() []CertificateFormat {
	return []CertificateFormat{Cer, Pfx}
}

// CertificateState enumerates the values for certificate state.
type CertificateState string

const (
	// Active The certificate is available for use in pools.
	Active CertificateState = "active"
	// DeleteFailed The user requested that the certificate be deleted, but there are pools that still have
	// references to the certificate, or it is still installed on one or more compute nodes. (The latter can
	// occur if the certificate has been removed from the pool, but the node has not yet restarted. Nodes
	// refresh their certificates only when they restart.) You may use the cancel certificate delete operation
	// to cancel the delete, or the delete certificate operation to retry the delete.
	DeleteFailed CertificateState = "deletefailed"
	// Deleting The user has requested that the certificate be deleted, but the delete operation has not yet
	// completed. You may not reference the certificate when creating or updating pools.
	Deleting CertificateState = "deleting"
)

// PossibleCertificateStateValues returns an array of possible values for the CertificateState const type.
func PossibleCertificateStateValues() []CertificateState {
	return []CertificateState{Active, DeleteFailed, Deleting}
}

// CertificateStoreLocation enumerates the values for certificate store location.
type CertificateStoreLocation string

const (
	// CurrentUser Certificates should be installed to the CurrentUser certificate store.
	CurrentUser CertificateStoreLocation = "currentuser"
	// LocalMachine Certificates should be installed to the LocalMachine certificate store.
	LocalMachine CertificateStoreLocation = "localmachine"
)

// PossibleCertificateStoreLocationValues returns an array of possible values for the CertificateStoreLocation const type.
func PossibleCertificateStoreLocationValues() []CertificateStoreLocation {
	return []CertificateStoreLocation{CurrentUser, LocalMachine}
}

// CertificateVisibility enumerates the values for certificate visibility.
type CertificateVisibility string

const (
	// CertificateVisibilityRemoteUser The certificate should be visible to the user accounts under which users
	// remotely access the node.
	CertificateVisibilityRemoteUser CertificateVisibility = "remoteuser"
	// CertificateVisibilityStartTask The certificate should be visible to the user account under which the
	// start task is run.
	CertificateVisibilityStartTask CertificateVisibility = "starttask"
	// CertificateVisibilityTask The certificate should be visible to the user accounts under which job tasks
	// are run.
	CertificateVisibilityTask CertificateVisibility = "task"
)

// PossibleCertificateVisibilityValues returns an array of possible values for the CertificateVisibility const type.
func PossibleCertificateVisibilityValues() []CertificateVisibility {
	return []CertificateVisibility{CertificateVisibilityRemoteUser, CertificateVisibilityStartTask, CertificateVisibilityTask}
}

// ComputeNodeDeallocationOption enumerates the values for compute node deallocation option.
type ComputeNodeDeallocationOption string

const (
	// Requeue Terminate running task processes and requeue the tasks. The tasks will run again when a node is
	// available. Remove nodes as soon as tasks have been terminated.
	Requeue ComputeNodeDeallocationOption = "requeue"
	// RetainedData Allow currently running tasks to complete, then wait for all task data retention periods to
	// expire. Schedule no new tasks while waiting. Remove nodes when all task retention periods have expired.
	RetainedData ComputeNodeDeallocationOption = "retaineddata"
	// TaskCompletion Allow currently running tasks to complete. Schedule no new tasks while waiting. Remove
	// nodes when all tasks have completed.
	TaskCompletion ComputeNodeDeallocationOption = "taskcompletion"
	// Terminate Terminate running tasks. The tasks will be completed with failureInfo indicating that they
	// were terminated, and will not run again. Remove nodes as soon as tasks have been terminated.
	Terminate ComputeNodeDeallocationOption = "terminate"
)

// PossibleComputeNodeDeallocationOptionValues returns an array of possible values for the ComputeNodeDeallocationOption const type.
func PossibleComputeNodeDeallocationOptionValues() []ComputeNodeDeallocationOption {
	return []ComputeNodeDeallocationOption{Requeue, RetainedData, TaskCompletion, Terminate}
}

// ComputeNodeFillType enumerates the values for compute node fill type.
type ComputeNodeFillType string

const (
	// Pack As many tasks as possible (maxTasksPerNode) should be assigned to each node in the pool before any
	// tasks are assigned to the next node in the pool.
	Pack ComputeNodeFillType = "pack"
	// Spread Tasks should be assigned evenly across all nodes in the pool.
	Spread ComputeNodeFillType = "spread"
)

// PossibleComputeNodeFillTypeValues returns an array of possible values for the ComputeNodeFillType const type.
func PossibleComputeNodeFillTypeValues() []ComputeNodeFillType {
	return []ComputeNodeFillType{Pack, Spread}
}

// ComputeNodeRebootOption enumerates the values for compute node reboot option.
type ComputeNodeRebootOption string

const (
	// ComputeNodeRebootOptionRequeue Terminate running task processes and requeue the tasks. The tasks will
	// run again when a node is available. Restart the node as soon as tasks have been terminated.
	ComputeNodeRebootOptionRequeue ComputeNodeRebootOption = "requeue"
	// ComputeNodeRebootOptionRetainedData Allow currently running tasks to complete, then wait for all task
	// data retention periods to expire. Schedule no new tasks while waiting. Restart the node when all task
	// retention periods have expired.
	ComputeNodeRebootOptionRetainedData ComputeNodeRebootOption = "retaineddata"
	// ComputeNodeRebootOptionTaskCompletion Allow currently running tasks to complete. Schedule no new tasks
	// while waiting. Restart the node when all tasks have completed.
	ComputeNodeRebootOptionTaskCompletion ComputeNodeRebootOption = "taskcompletion"
	// ComputeNodeRebootOptionTerminate Terminate running tasks. The tasks will be completed with failureInfo
	// indicating that they were terminated, and will not run again. Restart the node as soon as tasks have
	// been terminated.
	ComputeNodeRebootOptionTerminate ComputeNodeRebootOption = "terminate"
)

// PossibleComputeNodeRebootOptionValues returns an array of possible values for the ComputeNodeRebootOption const type.
func PossibleComputeNodeRebootOptionValues() []ComputeNodeRebootOption {
	return []ComputeNodeRebootOption{ComputeNodeRebootOptionRequeue, ComputeNodeRebootOptionRetainedData, ComputeNodeRebootOptionTaskCompletion, ComputeNodeRebootOptionTerminate}
}

// ComputeNodeReimageOption enumerates the values for compute node reimage option.
type ComputeNodeReimageOption string

const (
	// ComputeNodeReimageOptionRequeue Terminate running task processes and requeue the tasks. The tasks will
	// run again when a node is available. Reimage the node as soon as tasks have been terminated.
	ComputeNodeReimageOptionRequeue ComputeNodeReimageOption = "requeue"
	// ComputeNodeReimageOptionRetainedData Allow currently running tasks to complete, then wait for all task
	// data retention periods to expire. Schedule no new tasks while waiting. Reimage the node when all task
	// retention periods have expired.
	ComputeNodeReimageOptionRetainedData ComputeNodeReimageOption = "retaineddata"
	// ComputeNodeReimageOptionTaskCompletion Allow currently running tasks to complete. Schedule no new tasks
	// while waiting. Reimage the node when all tasks have completed.
	ComputeNodeReimageOptionTaskCompletion ComputeNodeReimageOption = "taskcompletion"
	// ComputeNodeReimageOptionTerminate Terminate running tasks. The tasks will be completed with failureInfo
	// indicating that they were terminated, and will not run again. Reimage the node as soon as tasks have
	// been terminated.
	ComputeNodeReimageOptionTerminate ComputeNodeReimageOption = "terminate"
)

// PossibleComputeNodeReimageOptionValues returns an array of possible values for the ComputeNodeReimageOption const type.
func PossibleComputeNodeReimageOptionValues() []ComputeNodeReimageOption {
	return []ComputeNodeReimageOption{ComputeNodeReimageOptionRequeue, ComputeNodeReimageOptionRetainedData, ComputeNodeReimageOptionTaskCompletion, ComputeNodeReimageOptionTerminate}
}

// ComputeNodeState enumerates the values for compute node state.
type ComputeNodeState string

const (
	// Creating The Batch service has obtained the underlying virtual machine from Azure Compute, but it has
	// not yet started to join the pool.
	Creating ComputeNodeState = "creating"
	// Idle The node is not currently running a task.
	Idle ComputeNodeState = "idle"
	// LeavingPool The node is leaving the pool, either because the user explicitly removed it or because the
	// pool is resizing or autoscaling down.
	LeavingPool ComputeNodeState = "leavingpool"
	// Offline The node is not currently running a task, and scheduling of new tasks to the node is disabled.
	Offline ComputeNodeState = "offline"
	// Preempted The low-priority node has been preempted. Tasks which were running on the node when it was
	// preempted will be rescheduled when another node becomes available.
	Preempted ComputeNodeState = "preempted"
	// Rebooting The node is rebooting.
	Rebooting ComputeNodeState = "rebooting"
	// Reimaging The node is reimaging.
	Reimaging ComputeNodeState = "reimaging"
	// Running The node is running one or more tasks (other than a start task).
	Running ComputeNodeState = "running"
	// Starting The Batch service is starting on the underlying virtual machine.
	Starting ComputeNodeState = "starting"
	// StartTaskFailed The start task has failed on the compute node (and exhausted all retries), and
	// waitForSuccess is set. The node is not usable for running tasks.
	StartTaskFailed ComputeNodeState = "starttaskfailed"
	// Unknown The Batch service has lost contact with the node, and does not know its true state.
	Unknown ComputeNodeState = "unknown"
	// Unusable The node cannot be used for task execution due to errors.
	Unusable ComputeNodeState = "unusable"
	// WaitingForStartTask The start task has started running on the compute node, but waitForSuccess is set
	// and the start task has not yet completed.
	WaitingForStartTask ComputeNodeState = "waitingforstarttask"
)

// PossibleComputeNodeStateValues returns an array of possible values for the ComputeNodeState const type.
func PossibleComputeNodeStateValues() []ComputeNodeState {
	return []ComputeNodeState{Creating, Idle, LeavingPool, Offline, Preempted, Rebooting, Reimaging, Running, Starting, StartTaskFailed, Unknown, Unusable, WaitingForStartTask}
}

// DependencyAction enumerates the values for dependency action.
type DependencyAction string

const (
	// Block Block the task's dependencies.
	Block DependencyAction = "block"
	// Satisfy Satisfy the task's dependencies.
	Satisfy DependencyAction = "satisfy"
)

// PossibleDependencyActionValues returns an array of possible values for the DependencyAction const type.
func PossibleDependencyActionValues() []DependencyAction {
	return []DependencyAction{Block, Satisfy}
}

// DisableComputeNodeSchedulingOption enumerates the values for disable compute node scheduling option.
type DisableComputeNodeSchedulingOption string

const (
	// DisableComputeNodeSchedulingOptionRequeue Terminate running task processes and requeue the tasks. The
	// tasks may run again on other compute nodes, or when task scheduling is re-enabled on this node. Enter
	// offline state as soon as tasks have been terminated.
	DisableComputeNodeSchedulingOptionRequeue DisableComputeNodeSchedulingOption = "requeue"
	// DisableComputeNodeSchedulingOptionTaskCompletion Allow currently running tasks to complete. Schedule no
	// new tasks while waiting. Enter offline state when all tasks have completed.
	DisableComputeNodeSchedulingOptionTaskCompletion DisableComputeNodeSchedulingOption = "taskcompletion"
	// DisableComputeNodeSchedulingOptionTerminate Terminate running tasks. The tasks will be completed with
	// failureInfo indicating that they were terminated, and will not run again. Enter offline state as soon as
	// tasks have been terminated.
	DisableComputeNodeSchedulingOptionTerminate DisableComputeNodeSchedulingOption = "terminate"
)

// PossibleDisableComputeNodeSchedulingOptionValues returns an array of possible values for the DisableComputeNodeSchedulingOption const type.
func PossibleDisableComputeNodeSchedulingOptionValues() []DisableComputeNodeSchedulingOption {
	return []DisableComputeNodeSchedulingOption{DisableComputeNodeSchedulingOptionRequeue, DisableComputeNodeSchedulingOptionTaskCompletion, DisableComputeNodeSchedulingOptionTerminate}
}

// DisableJobOption enumerates the values for disable job option.
type DisableJobOption string

const (
	// DisableJobOptionRequeue Terminate running tasks and requeue them. The tasks will run again when the job
	// is enabled.
	DisableJobOptionRequeue DisableJobOption = "requeue"
	// DisableJobOptionTerminate Terminate running tasks. The tasks will be completed with failureInfo
	// indicating that they were terminated, and will not run again.
	DisableJobOptionTerminate DisableJobOption = "terminate"
	// DisableJobOptionWait Allow currently running tasks to complete.
	DisableJobOptionWait DisableJobOption = "wait"
)

// PossibleDisableJobOptionValues returns an array of possible values for the DisableJobOption const type.
func PossibleDisableJobOptionValues() []DisableJobOption {
	return []DisableJobOption{DisableJobOptionRequeue, DisableJobOptionTerminate, DisableJobOptionWait}
}

// ElevationLevel enumerates the values for elevation level.
type ElevationLevel string

const (
	// Admin The user is a user with elevated access and operates with full Administrator permissions.
	Admin ElevationLevel = "admin"
	// NonAdmin The user is a standard user without elevated access.
	NonAdmin ElevationLevel = "nonadmin"
)

// PossibleElevationLevelValues returns an array of possible values for the ElevationLevel const type.
func PossibleElevationLevelValues() []ElevationLevel {
	return []ElevationLevel{Admin, NonAdmin}
}

// ErrorCategory enumerates the values for error category.
type ErrorCategory string

const (
	// ServerError The error is due to an internal server issue.
	ServerError ErrorCategory = "servererror"
	// UserError The error is due to a user issue, such as misconfiguration.
	UserError ErrorCategory = "usererror"
)

// PossibleErrorCategoryValues returns an array of possible values for the ErrorCategory const type.
func PossibleErrorCategoryValues() []ErrorCategory {
	return []ErrorCategory{ServerError, UserError}
}

// InboundEndpointProtocol enumerates the values for inbound endpoint protocol.
type InboundEndpointProtocol string

const (
	// TCP Use TCP for the endpoint.
	TCP InboundEndpointProtocol = "tcp"
	// UDP Use UDP for the endpoint.
	UDP InboundEndpointProtocol = "udp"
)

// PossibleInboundEndpointProtocolValues returns an array of possible values for the InboundEndpointProtocol const type.
func PossibleInboundEndpointProtocolValues() []InboundEndpointProtocol {
	return []InboundEndpointProtocol{TCP, UDP}
}

// JobAction enumerates the values for job action.
type JobAction string

const (
	// JobActionDisable Disable the job. This is equivalent to calling the disable job API, with a disableTasks
	// value of requeue.
	JobActionDisable JobAction = "disable"
	// JobActionNone Take no action.
	JobActionNone JobAction = "none"
	// JobActionTerminate Terminate the job. The terminateReason in the job's executionInfo is set to
	// "TaskFailed".
	JobActionTerminate JobAction = "terminate"
)

// PossibleJobActionValues returns an array of possible values for the JobAction const type.
func PossibleJobActionValues() []JobAction {
	return []JobAction{JobActionDisable, JobActionNone, JobActionTerminate}
}

// JobPreparationTaskState enumerates the values for job preparation task state.
type JobPreparationTaskState string

const (
	// JobPreparationTaskStateCompleted The task has exited with exit code 0, or the task has exhausted its
	// retry limit, or the Batch service was unable to start the task due to task preparation errors (such as
	// resource file download failures).
	JobPreparationTaskStateCompleted JobPreparationTaskState = "completed"
	// JobPreparationTaskStateRunning The task is currently running (including retrying).
	JobPreparationTaskStateRunning JobPreparationTaskState = "running"
)

// PossibleJobPreparationTaskStateValues returns an array of possible values for the JobPreparationTaskState const type.
func PossibleJobPreparationTaskStateValues() []JobPreparationTaskState {
	return []JobPreparationTaskState{JobPreparationTaskStateCompleted, JobPreparationTaskStateRunning}
}

// JobReleaseTaskState enumerates the values for job release task state.
type JobReleaseTaskState string

const (
	// JobReleaseTaskStateCompleted The task has exited with exit code 0, or the task has exhausted its retry
	// limit, or the Batch service was unable to start the task due to task preparation errors (such as
	// resource file download failures).
	JobReleaseTaskStateCompleted JobReleaseTaskState = "completed"
	// JobReleaseTaskStateRunning The task is currently running (including retrying).
	JobReleaseTaskStateRunning JobReleaseTaskState = "running"
)

// PossibleJobReleaseTaskStateValues returns an array of possible values for the JobReleaseTaskState const type.
func PossibleJobReleaseTaskStateValues() []JobReleaseTaskState {
	return []JobReleaseTaskState{JobReleaseTaskStateCompleted, JobReleaseTaskStateRunning}
}

// JobScheduleState enumerates the values for job schedule state.
type JobScheduleState string

const (
	// JobScheduleStateActive The job schedule is active and will create jobs as per its schedule.
	JobScheduleStateActive JobScheduleState = "active"
	// JobScheduleStateCompleted The schedule has terminated, either by reaching its end time or by the user
	// terminating it explicitly.
	JobScheduleStateCompleted JobScheduleState = "completed"
	// JobScheduleStateDeleting The user has requested that the schedule be deleted, but the delete operation
	// is still in progress. The scheduler will not initiate any new jobs for this schedule, and will delete
	// any existing jobs and tasks under the schedule, including any active job. The schedule will be deleted
	// when all jobs and tasks under the schedule have been deleted.
	JobScheduleStateDeleting JobScheduleState = "deleting"
	// JobScheduleStateDisabled The user has disabled the schedule. The scheduler will not initiate any new
	// jobs will on this schedule, but any existing active job will continue to run.
	JobScheduleStateDisabled JobScheduleState = "disabled"
	// JobScheduleStateTerminating The schedule has no more work to do, or has been explicitly terminated by
	// the user, but the termination operation is still in progress. The scheduler will not initiate any new
	// jobs for this schedule, nor is any existing job active.
	JobScheduleStateTerminating JobScheduleState = "terminating"
)

// PossibleJobScheduleStateValues returns an array of possible values for the JobScheduleState const type.
func PossibleJobScheduleStateValues() []JobScheduleState {
	return []JobScheduleState{JobScheduleStateActive, JobScheduleStateCompleted, JobScheduleStateDeleting, JobScheduleStateDisabled, JobScheduleStateTerminating}
}

// JobState enumerates the values for job state.
type JobState string

const (
	// JobStateActive The job is available to have tasks scheduled.
	JobStateActive JobState = "active"
	// JobStateCompleted All tasks have terminated, and the system will not accept any more tasks or any
	// further changes to the job.
	JobStateCompleted JobState = "completed"
	// JobStateDeleting A user has requested that the job be deleted, but the delete operation is still in
	// progress (for example, because the system is still terminating running tasks).
	JobStateDeleting JobState = "deleting"
	// JobStateDisabled A user has disabled the job. No tasks are running, and no new tasks will be scheduled.
	JobStateDisabled JobState = "disabled"
	// JobStateDisabling A user has requested that the job be disabled, but the disable operation is still in
	// progress (for example, waiting for tasks to terminate).
	JobStateDisabling JobState = "disabling"
	// JobStateEnabling A user has requested that the job be enabled, but the enable operation is still in
	// progress.
	JobStateEnabling JobState = "enabling"
	// JobStateTerminating The job is about to complete, either because a Job Manager task has completed or
	// because the user has terminated the job, but the terminate operation is still in progress (for example,
	// because Job Release tasks are running).
	JobStateTerminating JobState = "terminating"
)

// PossibleJobStateValues returns an array of possible values for the JobState const type.
func PossibleJobStateValues() []JobState {
	return []JobState{JobStateActive, JobStateCompleted, JobStateDeleting, JobStateDisabled, JobStateDisabling, JobStateEnabling, JobStateTerminating}
}

// NetworkSecurityGroupRuleAccess enumerates the values for network security group rule access.
type NetworkSecurityGroupRuleAccess string

const (
	// Allow Allow access.
	Allow NetworkSecurityGroupRuleAccess = "allow"
	// Deny Deny access.
	Deny NetworkSecurityGroupRuleAccess = "deny"
)

// PossibleNetworkSecurityGroupRuleAccessValues returns an array of possible values for the NetworkSecurityGroupRuleAccess const type.
func PossibleNetworkSecurityGroupRuleAccessValues() []NetworkSecurityGroupRuleAccess {
	return []NetworkSecurityGroupRuleAccess{Allow, Deny}
}

// OnAllTasksComplete enumerates the values for on all tasks complete.
type OnAllTasksComplete string

const (
	// NoAction Do nothing. The job remains active unless terminated or disabled by some other means.
	NoAction OnAllTasksComplete = "noaction"
	// TerminateJob Terminate the job. The job's terminateReason is set to 'AllTasksComplete'.
	TerminateJob OnAllTasksComplete = "terminatejob"
)

// PossibleOnAllTasksCompleteValues returns an array of possible values for the OnAllTasksComplete const type.
func PossibleOnAllTasksCompleteValues() []OnAllTasksComplete {
	return []OnAllTasksComplete{NoAction, TerminateJob}
}

// OnTaskFailure enumerates the values for on task failure.
type OnTaskFailure string

const (
	// OnTaskFailureNoAction Do nothing. The job remains active unless terminated or disabled by some other
	// means.
	OnTaskFailureNoAction OnTaskFailure = "noaction"
	// OnTaskFailurePerformExitOptionsJobAction Take the action associated with the task exit condition in the
	// task's exitConditions collection. (This may still result in no action being taken, if that is what the
	// task specifies.)
	OnTaskFailurePerformExitOptionsJobAction OnTaskFailure = "performexitoptionsjobaction"
)

// PossibleOnTaskFailureValues returns an array of possible values for the OnTaskFailure const type.
func PossibleOnTaskFailureValues() []OnTaskFailure {
	return []OnTaskFailure{OnTaskFailureNoAction, OnTaskFailurePerformExitOptionsJobAction}
}

// OSType enumerates the values for os type.
type OSType string

const (
	// Linux The Linux operating system.
	Linux OSType = "linux"
	// Windows The Windows operating system.
	Windows OSType = "windows"
)

// PossibleOSTypeValues returns an array of possible values for the OSType const type.
func PossibleOSTypeValues() []OSType {
	return []OSType{Linux, Windows}
}

// OutputFileUploadCondition enumerates the values for output file upload condition.
type OutputFileUploadCondition string

const (
	// OutputFileUploadConditionTaskCompletion Upload the file(s) after the task process exits, no matter what
	// the exit code was.
	OutputFileUploadConditionTaskCompletion OutputFileUploadCondition = "taskcompletion"
	// OutputFileUploadConditionTaskFailure Upload the file(s) only after the task process exits with a nonzero
	// exit code.
	OutputFileUploadConditionTaskFailure OutputFileUploadCondition = "taskfailure"
	// OutputFileUploadConditionTaskSuccess Upload the file(s) only after the task process exits with an exit
	// code of 0.
	OutputFileUploadConditionTaskSuccess OutputFileUploadCondition = "tasksuccess"
)

// PossibleOutputFileUploadConditionValues returns an array of possible values for the OutputFileUploadCondition const type.
func PossibleOutputFileUploadConditionValues() []OutputFileUploadCondition {
	return []OutputFileUploadCondition{OutputFileUploadConditionTaskCompletion, OutputFileUploadConditionTaskFailure, OutputFileUploadConditionTaskSuccess}
}

// PoolLifetimeOption enumerates the values for pool lifetime option.
type PoolLifetimeOption string

const (
	// PoolLifetimeOptionJob The pool exists for the lifetime of the job to which it is dedicated. The Batch
	// service creates the pool when it creates the job. If the 'job' option is applied to a job schedule, the
	// Batch service creates a new auto pool for every job created on the schedule.
	PoolLifetimeOptionJob PoolLifetimeOption = "job"
	// PoolLifetimeOptionJobSchedule The pool exists for the lifetime of the job schedule. The Batch Service
	// creates the pool when it creates the first job on the schedule. You may apply this option only to job
	// schedules, not to jobs.
	PoolLifetimeOptionJobSchedule PoolLifetimeOption = "jobschedule"
)

// PossiblePoolLifetimeOptionValues returns an array of possible values for the PoolLifetimeOption const type.
func PossiblePoolLifetimeOptionValues() []PoolLifetimeOption {
	return []PoolLifetimeOption{PoolLifetimeOptionJob, PoolLifetimeOptionJobSchedule}
}

// PoolState enumerates the values for pool state.
type PoolState string

const (
	// PoolStateActive The pool is available to run tasks subject to the availability of compute nodes.
	PoolStateActive PoolState = "active"
	// PoolStateDeleting The user has requested that the pool be deleted, but the delete operation has not yet
	// completed.
	PoolStateDeleting PoolState = "deleting"
	// PoolStateUpgrading The user has requested that the operating system of the pool's nodes be upgraded, but
	// the upgrade operation has not yet completed (that is, some nodes in the pool have not yet been
	// upgraded). While upgrading, the pool may be able to run tasks (with reduced capacity) but this is not
	// guaranteed.
	PoolStateUpgrading PoolState = "upgrading"
)

// PossiblePoolStateValues returns an array of possible values for the PoolState const type.
func PossiblePoolStateValues() []PoolState {
	return []PoolState{PoolStateActive, PoolStateDeleting, PoolStateUpgrading}
}

// SchedulingState enumerates the values for scheduling state.
type SchedulingState string

const (
	// Disabled No new tasks will be scheduled on the node. Tasks already running on the node may still run to
	// completion. All nodes start with scheduling enabled.
	Disabled SchedulingState = "disabled"
	// Enabled Tasks can be scheduled on the node.
	Enabled SchedulingState = "enabled"
)

// PossibleSchedulingStateValues returns an array of possible values for the SchedulingState const type.
func PossibleSchedulingStateValues() []SchedulingState {
	return []SchedulingState{Disabled, Enabled}
}

// StartTaskState enumerates the values for start task state.
type StartTaskState string

const (
	// StartTaskStateCompleted The start task has exited with exit code 0, or the start task has failed and the
	// retry limit has reached, or the start task process did not run due to task preparation errors (such as
	// resource file download failures).
	StartTaskStateCompleted StartTaskState = "completed"
	// StartTaskStateRunning The start task is currently running.
	StartTaskStateRunning StartTaskState = "running"
)

// PossibleStartTaskStateValues returns an array of possible values for the StartTaskState const type.
func PossibleStartTaskStateValues() []StartTaskState {
	return []StartTaskState{StartTaskStateCompleted, StartTaskStateRunning}
}

// StorageAccountType enumerates the values for storage account type.
type StorageAccountType string

const (
	// PremiumLRS The data disk should use premium locally redundant storage.
	PremiumLRS StorageAccountType = "premium_lrs"
	// StandardLRS The data disk should use standard locally redundant storage.
	StandardLRS StorageAccountType = "standard_lrs"
)

// PossibleStorageAccountTypeValues returns an array of possible values for the StorageAccountType const type.
func PossibleStorageAccountTypeValues() []StorageAccountType {
	return []StorageAccountType{PremiumLRS, StandardLRS}
}

// SubtaskState enumerates the values for subtask state.
type SubtaskState string

const (
	// SubtaskStateCompleted The task is no longer eligible to run, usually because the task has finished
	// successfully, or the task has finished unsuccessfully and has exhausted its retry limit. A task is also
	// marked as completed if an error occurred launching the task, or when the task has been terminated.
	SubtaskStateCompleted SubtaskState = "completed"
	// SubtaskStatePreparing The task has been assigned to a compute node, but is waiting for a required Job
	// Preparation task to complete on the node. If the Job Preparation task succeeds, the task will move to
	// running. If the Job Preparation task fails, the task will return to active and will be eligible to be
	// assigned to a different node.
	SubtaskStatePreparing SubtaskState = "preparing"
	// SubtaskStateRunning The task is running on a compute node. This includes task-level preparation such as
	// downloading resource files or deploying application packages specified on the task - it does not
	// necessarily mean that the task command line has started executing.
	SubtaskStateRunning SubtaskState = "running"
)

// PossibleSubtaskStateValues returns an array of possible values for the SubtaskState const type.
func PossibleSubtaskStateValues() []SubtaskState {
	return []SubtaskState{SubtaskStateCompleted, SubtaskStatePreparing, SubtaskStateRunning}
}

// TaskAddStatus enumerates the values for task add status.
type TaskAddStatus string

const (
	// TaskAddStatusClientError The task failed to add due to a client error and should not be retried without
	// modifying the request as appropriate.
	TaskAddStatusClientError TaskAddStatus = "clienterror"
	// TaskAddStatusServerError Task failed to add due to a server error and can be retried without
	// modification.
	TaskAddStatusServerError TaskAddStatus = "servererror"
	// TaskAddStatusSuccess The task was added successfully.
	TaskAddStatusSuccess TaskAddStatus = "success"
)

// PossibleTaskAddStatusValues returns an array of possible values for the TaskAddStatus const type.
func PossibleTaskAddStatusValues() []TaskAddStatus {
	return []TaskAddStatus{TaskAddStatusClientError, TaskAddStatusServerError, TaskAddStatusSuccess}
}

// TaskExecutionResult enumerates the values for task execution result.
type TaskExecutionResult string

const (
	// Failure There was an error during processing of the task. The failure may have occurred before the task
	// process was launched, while the task process was executing, or after the task process exited.
	Failure TaskExecutionResult = "failure"
	// Success The task ran successfully.
	Success TaskExecutionResult = "success"
)

// PossibleTaskExecutionResultValues returns an array of possible values for the TaskExecutionResult const type.
func PossibleTaskExecutionResultValues() []TaskExecutionResult {
	return []TaskExecutionResult{Failure, Success}
}

// TaskState enumerates the values for task state.
type TaskState string

const (
	// TaskStateActive The task is queued and able to run, but is not currently assigned to a compute node. A
	// task enters this state when it is created, when it is enabled after being disabled, or when it is
	// awaiting a retry after a failed run.
	TaskStateActive TaskState = "active"
	// TaskStateCompleted The task is no longer eligible to run, usually because the task has finished
	// successfully, or the task has finished unsuccessfully and has exhausted its retry limit. A task is also
	// marked as completed if an error occurred launching the task, or when the task has been terminated.
	TaskStateCompleted TaskState = "completed"
	// TaskStatePreparing The task has been assigned to a compute node, but is waiting for a required Job
	// Preparation task to complete on the node. If the Job Preparation task succeeds, the task will move to
	// running. If the Job Preparation task fails, the task will return to active and will be eligible to be
	// assigned to a different node.
	TaskStatePreparing TaskState = "preparing"
	// TaskStateRunning The task is running on a compute node. This includes task-level preparation such as
	// downloading resource files or deploying application packages specified on the task - it does not
	// necessarily mean that the task command line has started executing.
	TaskStateRunning TaskState = "running"
)

// PossibleTaskStateValues returns an array of possible values for the TaskState const type.
func PossibleTaskStateValues() []TaskState {
	return []TaskState{TaskStateActive, TaskStateCompleted, TaskStatePreparing, TaskStateRunning}
}
