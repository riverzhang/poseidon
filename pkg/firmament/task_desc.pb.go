/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: task_desc.proto

package firmament

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TaskDescriptor_TaskState int32

const (
	TaskDescriptor_CREATED   TaskDescriptor_TaskState = 0
	TaskDescriptor_BLOCKING  TaskDescriptor_TaskState = 1
	TaskDescriptor_RUNNABLE  TaskDescriptor_TaskState = 2
	TaskDescriptor_ASSIGNED  TaskDescriptor_TaskState = 3
	TaskDescriptor_RUNNING   TaskDescriptor_TaskState = 4
	TaskDescriptor_COMPLETED TaskDescriptor_TaskState = 5
	TaskDescriptor_FAILED    TaskDescriptor_TaskState = 6
	TaskDescriptor_ABORTED   TaskDescriptor_TaskState = 7
	TaskDescriptor_DELEGATED TaskDescriptor_TaskState = 8
	TaskDescriptor_UNKNOWN   TaskDescriptor_TaskState = 9
)

var TaskDescriptor_TaskState_name = map[int32]string{
	0: "CREATED",
	1: "BLOCKING",
	2: "RUNNABLE",
	3: "ASSIGNED",
	4: "RUNNING",
	5: "COMPLETED",
	6: "FAILED",
	7: "ABORTED",
	8: "DELEGATED",
	9: "UNKNOWN",
}

var TaskDescriptor_TaskState_value = map[string]int32{
	"CREATED":   0,
	"BLOCKING":  1,
	"RUNNABLE":  2,
	"ASSIGNED":  3,
	"RUNNING":   4,
	"COMPLETED": 5,
	"FAILED":    6,
	"ABORTED":   7,
	"DELEGATED": 8,
	"UNKNOWN":   9,
}

func (x TaskDescriptor_TaskState) String() string {
	return proto.EnumName(TaskDescriptor_TaskState_name, int32(x))
}

func (TaskDescriptor_TaskState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_37c54fe0f119fae3, []int{0, 0}
}

type TaskDescriptor_TaskType int32

const (
	TaskDescriptor_SHEEP  TaskDescriptor_TaskType = 0
	TaskDescriptor_RABBIT TaskDescriptor_TaskType = 1
	TaskDescriptor_DEVIL  TaskDescriptor_TaskType = 2
	TaskDescriptor_TURTLE TaskDescriptor_TaskType = 3
)

var TaskDescriptor_TaskType_name = map[int32]string{
	0: "SHEEP",
	1: "RABBIT",
	2: "DEVIL",
	3: "TURTLE",
}

var TaskDescriptor_TaskType_value = map[string]int32{
	"SHEEP":  0,
	"RABBIT": 1,
	"DEVIL":  2,
	"TURTLE": 3,
}

func (x TaskDescriptor_TaskType) String() string {
	return proto.EnumName(TaskDescriptor_TaskType_name, int32(x))
}

func (TaskDescriptor_TaskType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_37c54fe0f119fae3, []int{0, 1}
}

// TaskDescriptor describes a task in firmament scheduler.
type TaskDescriptor struct {
	Uid   uint64                   `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name  string                   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	State TaskDescriptor_TaskState `protobuf:"varint,3,opt,name=state,proto3,enum=firmament.TaskDescriptor_TaskState" json:"state,omitempty"`
	JobId string                   `protobuf:"bytes,4,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	Index uint64                   `protobuf:"varint,5,opt,name=index,proto3" json:"index,omitempty"`
	// Inputs/outputs
	Dependencies []*ReferenceDescriptor `protobuf:"bytes,6,rep,name=dependencies,proto3" json:"dependencies,omitempty"`
	Outputs      []*ReferenceDescriptor `protobuf:"bytes,7,rep,name=outputs,proto3" json:"outputs,omitempty"`
	// Command and arguments
	Binary string   `protobuf:"bytes,8,opt,name=binary,proto3" json:"binary,omitempty"`
	Args   []string `protobuf:"bytes,9,rep,name=args,proto3" json:"args,omitempty"`
	// spawned is the list of children tasks
	Spawned []*TaskDescriptor `protobuf:"bytes,10,rep,name=spawned,proto3" json:"spawned,omitempty"`
	// Runtime meta-data
	ScheduledToResource   string `protobuf:"bytes,11,opt,name=scheduled_to_resource,json=scheduledToResource,proto3" json:"scheduled_to_resource,omitempty"`
	LastHeartbeatLocation string `protobuf:"bytes,12,opt,name=last_heartbeat_location,json=lastHeartbeatLocation,proto3" json:"last_heartbeat_location,omitempty"`
	LastHeartbeatTime     uint64 `protobuf:"varint,13,opt,name=last_heartbeat_time,json=lastHeartbeatTime,proto3" json:"last_heartbeat_time,omitempty"`
	// Delegation info
	DelegatedTo   string `protobuf:"bytes,14,opt,name=delegated_to,json=delegatedTo,proto3" json:"delegated_to,omitempty"`
	DelegatedFrom string `protobuf:"bytes,15,opt,name=delegated_from,json=delegatedFrom,proto3" json:"delegated_from,omitempty"`
	// Timestamps
	SubmitTime uint64 `protobuf:"varint,16,opt,name=submit_time,json=submitTime,proto3" json:"submit_time,omitempty"`
	StartTime  uint64 `protobuf:"varint,17,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	FinishTime uint64 `protobuf:"varint,18,opt,name=finish_time,json=finishTime,proto3" json:"finish_time,omitempty"`
	// The total time spent unscheduled before previous runs.
	TotalUnscheduledTime uint64 `protobuf:"varint,19,opt,name=total_unscheduled_time,json=totalUnscheduledTime,proto3" json:"total_unscheduled_time,omitempty"`
	// The total time spent in previous runs. This field only gets updated when
	//   the task finishes running.
	TotalRunTime uint64 `protobuf:"varint,20,opt,name=total_run_time,json=totalRunTime,proto3" json:"total_run_time,omitempty"`
	// Deadlines
	RelativeDeadline uint64 `protobuf:"varint,21,opt,name=relative_deadline,json=relativeDeadline,proto3" json:"relative_deadline,omitempty"`
	AbsoluteDeadline uint64 `protobuf:"varint,22,opt,name=absolute_deadline,json=absoluteDeadline,proto3" json:"absolute_deadline,omitempty"`
	// Application-specific fields
	// TODO(malte): move these to sub-messages
	Port      uint64 `protobuf:"varint,23,opt,name=port,proto3" json:"port,omitempty"`
	InputSize uint64 `protobuf:"varint,24,opt,name=input_size,json=inputSize,proto3" json:"input_size,omitempty"`
	// TaskLib related stuff
	InjectTaskLib bool `protobuf:"varint,25,opt,name=inject_task_lib,json=injectTaskLib,proto3" json:"inject_task_lib,omitempty"`
	// Task resource request and priority
	ResourceRequest *ResourceVector `protobuf:"bytes,26,opt,name=resource_request,json=resourceRequest,proto3" json:"resource_request,omitempty"`
	Priority        uint32          `protobuf:"varint,27,opt,name=priority,proto3" json:"priority,omitempty"`
	// TODO(malte): move this to a policy-specific sub-message
	TaskType TaskDescriptor_TaskType `protobuf:"varint,28,opt,name=task_type,json=taskType,proto3,enum=firmament.TaskDescriptor_TaskType" json:"task_type,omitempty"`
	// Final report of a task after successful execution
	FinalReport *TaskFinalReport `protobuf:"bytes,29,opt,name=final_report,json=finalReport,proto3" json:"final_report,omitempty"`
	// Simulation related fields
	TraceJobId  uint64 `protobuf:"varint,30,opt,name=trace_job_id,json=traceJobId,proto3" json:"trace_job_id,omitempty"`
	TraceTaskId uint64 `protobuf:"varint,31,opt,name=trace_task_id,json=traceTaskId,proto3" json:"trace_task_id,omitempty"`
	// Task labels
	Labels []*Label `protobuf:"bytes,32,rep,name=labels,proto3" json:"labels,omitempty"`
	// Resource label selectors
	LabelSelectors []*LabelSelector `protobuf:"bytes,33,rep,name=label_selectors,json=labelSelectors,proto3" json:"label_selectors,omitempty"`
	//Affinity
	Affinity *Affinity `protobuf:"bytes,34,opt,name=affinity,proto3" json:"affinity,omitempty"`
	//NameSpace
	Namespace string `protobuf:"bytes,35,opt,name=namespace,proto3" json:"namespace,omitempty"`
	//Toleration
	Toleration           []*Toleration `protobuf:"bytes,36,rep,name=toleration,proto3" json:"toleration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *TaskDescriptor) Reset()         { *m = TaskDescriptor{} }
func (m *TaskDescriptor) String() string { return proto.CompactTextString(m) }
func (*TaskDescriptor) ProtoMessage()    {}
func (*TaskDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_37c54fe0f119fae3, []int{0}
}

func (m *TaskDescriptor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskDescriptor.Unmarshal(m, b)
}
func (m *TaskDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskDescriptor.Marshal(b, m, deterministic)
}
func (m *TaskDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskDescriptor.Merge(m, src)
}
func (m *TaskDescriptor) XXX_Size() int {
	return xxx_messageInfo_TaskDescriptor.Size(m)
}
func (m *TaskDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_TaskDescriptor proto.InternalMessageInfo

func (m *TaskDescriptor) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *TaskDescriptor) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TaskDescriptor) GetState() TaskDescriptor_TaskState {
	if m != nil {
		return m.State
	}
	return TaskDescriptor_CREATED
}

func (m *TaskDescriptor) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

func (m *TaskDescriptor) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *TaskDescriptor) GetDependencies() []*ReferenceDescriptor {
	if m != nil {
		return m.Dependencies
	}
	return nil
}

func (m *TaskDescriptor) GetOutputs() []*ReferenceDescriptor {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func (m *TaskDescriptor) GetBinary() string {
	if m != nil {
		return m.Binary
	}
	return ""
}

func (m *TaskDescriptor) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *TaskDescriptor) GetSpawned() []*TaskDescriptor {
	if m != nil {
		return m.Spawned
	}
	return nil
}

func (m *TaskDescriptor) GetScheduledToResource() string {
	if m != nil {
		return m.ScheduledToResource
	}
	return ""
}

func (m *TaskDescriptor) GetLastHeartbeatLocation() string {
	if m != nil {
		return m.LastHeartbeatLocation
	}
	return ""
}

func (m *TaskDescriptor) GetLastHeartbeatTime() uint64 {
	if m != nil {
		return m.LastHeartbeatTime
	}
	return 0
}

func (m *TaskDescriptor) GetDelegatedTo() string {
	if m != nil {
		return m.DelegatedTo
	}
	return ""
}

func (m *TaskDescriptor) GetDelegatedFrom() string {
	if m != nil {
		return m.DelegatedFrom
	}
	return ""
}

func (m *TaskDescriptor) GetSubmitTime() uint64 {
	if m != nil {
		return m.SubmitTime
	}
	return 0
}

func (m *TaskDescriptor) GetStartTime() uint64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *TaskDescriptor) GetFinishTime() uint64 {
	if m != nil {
		return m.FinishTime
	}
	return 0
}

func (m *TaskDescriptor) GetTotalUnscheduledTime() uint64 {
	if m != nil {
		return m.TotalUnscheduledTime
	}
	return 0
}

func (m *TaskDescriptor) GetTotalRunTime() uint64 {
	if m != nil {
		return m.TotalRunTime
	}
	return 0
}

func (m *TaskDescriptor) GetRelativeDeadline() uint64 {
	if m != nil {
		return m.RelativeDeadline
	}
	return 0
}

func (m *TaskDescriptor) GetAbsoluteDeadline() uint64 {
	if m != nil {
		return m.AbsoluteDeadline
	}
	return 0
}

func (m *TaskDescriptor) GetPort() uint64 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *TaskDescriptor) GetInputSize() uint64 {
	if m != nil {
		return m.InputSize
	}
	return 0
}

func (m *TaskDescriptor) GetInjectTaskLib() bool {
	if m != nil {
		return m.InjectTaskLib
	}
	return false
}

func (m *TaskDescriptor) GetResourceRequest() *ResourceVector {
	if m != nil {
		return m.ResourceRequest
	}
	return nil
}

func (m *TaskDescriptor) GetPriority() uint32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *TaskDescriptor) GetTaskType() TaskDescriptor_TaskType {
	if m != nil {
		return m.TaskType
	}
	return TaskDescriptor_SHEEP
}

func (m *TaskDescriptor) GetFinalReport() *TaskFinalReport {
	if m != nil {
		return m.FinalReport
	}
	return nil
}

func (m *TaskDescriptor) GetTraceJobId() uint64 {
	if m != nil {
		return m.TraceJobId
	}
	return 0
}

func (m *TaskDescriptor) GetTraceTaskId() uint64 {
	if m != nil {
		return m.TraceTaskId
	}
	return 0
}

func (m *TaskDescriptor) GetLabels() []*Label {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *TaskDescriptor) GetLabelSelectors() []*LabelSelector {
	if m != nil {
		return m.LabelSelectors
	}
	return nil
}

func (m *TaskDescriptor) GetAffinity() *Affinity {
	if m != nil {
		return m.Affinity
	}
	return nil
}

func (m *TaskDescriptor) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *TaskDescriptor) GetToleration() []*Toleration {
	if m != nil {
		return m.Toleration
	}
	return nil
}

func init() {
	proto.RegisterEnum("firmament.TaskDescriptor_TaskState", TaskDescriptor_TaskState_name, TaskDescriptor_TaskState_value)
	proto.RegisterEnum("firmament.TaskDescriptor_TaskType", TaskDescriptor_TaskType_name, TaskDescriptor_TaskType_value)
	proto.RegisterType((*TaskDescriptor)(nil), "firmament.TaskDescriptor")
}

func init() { proto.RegisterFile("task_desc.proto", fileDescriptor_37c54fe0f119fae3) }

var fileDescriptor_37c54fe0f119fae3 = []byte{
	// 980 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0x6d, 0x6f, 0x1a, 0xc7,
	0x13, 0x0f, 0xc1, 0x60, 0x18, 0x1e, 0x7c, 0x5e, 0x1b, 0x7b, 0xc3, 0x3f, 0x89, 0x09, 0xc9, 0xbf,
	0x42, 0xaa, 0xe4, 0x4a, 0x4e, 0x5b, 0xa5, 0x2f, 0xaa, 0x0a, 0xcc, 0xd9, 0xa1, 0xa1, 0x38, 0x5a,
	0x70, 0xfa, 0xf2, 0xb4, 0x70, 0x4b, 0xbc, 0xce, 0x71, 0x47, 0x77, 0xf7, 0xd2, 0x3a, 0x5f, 0xa3,
	0x5f, 0xa8, 0x1f, 0xad, 0xda, 0xb9, 0x3b, 0x38, 0x5b, 0x6a, 0xd5, 0x77, 0x3b, 0xbf, 0x87, 0x99,
	0x65, 0x6e, 0x76, 0x80, 0x3d, 0xc3, 0xf5, 0x27, 0xcf, 0x17, 0x7a, 0x71, 0xba, 0x56, 0x91, 0x89,
	0x48, 0x75, 0x29, 0xd5, 0x8a, 0xaf, 0x44, 0x68, 0xda, 0xb5, 0x80, 0xcf, 0x45, 0x90, 0xe0, 0xed,
	0x43, 0x0c, 0x3c, 0x2d, 0x02, 0xb1, 0x30, 0x91, 0xca, 0x50, 0x25, 0x96, 0x42, 0x89, 0x70, 0x21,
	0x72, 0x39, 0xda, 0x2d, 0x25, 0x74, 0x14, 0xab, 0x85, 0xf0, 0x3e, 0xe7, 0xc5, 0xc7, 0x58, 0x6b,
	0x29, 0x43, 0x1e, 0x78, 0x4a, 0xac, 0x23, 0x65, 0x52, 0xa2, 0xc9, 0x97, 0x4b, 0x19, 0x4a, 0x73,
	0x97, 0xc6, 0xfb, 0x26, 0x0a, 0x84, 0xe2, 0x46, 0x46, 0xa1, 0x4e, 0xa0, 0xee, 0x5f, 0x0d, 0x68,
	0xce, 0xb8, 0xfe, 0x34, 0x14, 0x7a, 0xa1, 0xe4, 0xda, 0x44, 0x8a, 0x38, 0x50, 0x8c, 0xa5, 0x4f,
	0x0b, 0x9d, 0x42, 0x6f, 0x87, 0xd9, 0x23, 0x21, 0xb0, 0x13, 0xf2, 0x95, 0xa0, 0x8f, 0x3b, 0x85,
	0x5e, 0x95, 0xe1, 0x99, 0xfc, 0x00, 0x25, 0x6d, 0xb8, 0x11, 0xb4, 0xd8, 0x29, 0xf4, 0x9a, 0x67,
	0x2f, 0x4f, 0x37, 0xbf, 0xef, 0xf4, 0x7e, 0x3e, 0x0c, 0xa7, 0x56, 0xca, 0x12, 0x07, 0x69, 0x41,
	0xf9, 0x36, 0x9a, 0x7b, 0xd2, 0xa7, 0x3b, 0x98, 0xb0, 0x74, 0x1b, 0xcd, 0x47, 0x3e, 0x39, 0x84,
	0x92, 0x0c, 0x7d, 0xf1, 0x07, 0x2d, 0x61, 0xe5, 0x24, 0x20, 0x03, 0xa8, 0xfb, 0x62, 0x2d, 0x42,
	0x5f, 0x84, 0x0b, 0x29, 0x34, 0x2d, 0x77, 0x8a, 0xbd, 0xda, 0xd9, 0xf3, 0x5c, 0x39, 0x96, 0xb5,
	0x6a, 0x5b, 0x93, 0xdd, 0xf3, 0x90, 0x37, 0xb0, 0x1b, 0xc5, 0x66, 0x1d, 0x1b, 0x4d, 0x77, 0xff,
	0x93, 0x3d, 0x93, 0x93, 0x23, 0x28, 0xcf, 0x65, 0xc8, 0xd5, 0x1d, 0xad, 0xe0, 0x55, 0xd3, 0xc8,
	0x76, 0x84, 0xab, 0x8f, 0x9a, 0x56, 0x3b, 0x45, 0xdb, 0x11, 0x7b, 0x26, 0xaf, 0x61, 0x57, 0xaf,
	0xf9, 0xef, 0xa1, 0xf0, 0x29, 0x60, 0x95, 0x27, 0xff, 0xd8, 0x13, 0x96, 0x29, 0xc9, 0x19, 0xb4,
	0xf4, 0xe2, 0x46, 0xf8, 0x71, 0x20, 0x7c, 0xcf, 0x44, 0x5e, 0xf6, 0x85, 0x69, 0x0d, 0xeb, 0x1d,
	0x6c, 0xc8, 0x59, 0xc4, 0x52, 0x8a, 0x7c, 0x0f, 0xc7, 0x01, 0xd7, 0xc6, 0xbb, 0x11, 0x5c, 0x99,
	0xb9, 0xe0, 0xc6, 0x0b, 0xa2, 0x05, 0x7e, 0x55, 0x5a, 0x47, 0x57, 0xcb, 0xd2, 0x6f, 0x33, 0x76,
	0x9c, 0x92, 0xe4, 0x14, 0x0e, 0x1e, 0xf8, 0x8c, 0x5c, 0x09, 0xda, 0xc0, 0x76, 0xef, 0xdf, 0xf3,
	0xcc, 0xe4, 0x4a, 0x90, 0x17, 0xb6, 0xf5, 0x81, 0xf8, 0xc8, 0x0d, 0xde, 0x8d, 0x36, 0x31, 0x79,
	0x6d, 0x83, 0xcd, 0x22, 0xf2, 0x7f, 0x68, 0x6e, 0x25, 0x4b, 0x15, 0xad, 0xe8, 0x1e, 0x8a, 0x1a,
	0x1b, 0xf4, 0x42, 0x45, 0x2b, 0x72, 0x02, 0x35, 0x1d, 0xcf, 0x57, 0x32, 0xad, 0xe8, 0x60, 0x45,
	0x48, 0x20, 0x2c, 0xf5, 0x0c, 0x40, 0x1b, 0xae, 0x52, 0x7e, 0x1f, 0xf9, 0x2a, 0x22, 0x48, 0x9f,
	0x40, 0xcd, 0x0e, 0xb2, 0xbe, 0x49, 0x78, 0x92, 0xf8, 0x13, 0x08, 0x05, 0xdf, 0xc2, 0x91, 0x89,
	0x0c, 0x0f, 0xbc, 0x38, 0xcc, 0xb5, 0xd3, 0x6a, 0x0f, 0x50, 0x7b, 0x88, 0xec, 0xf5, 0x96, 0x44,
	0xd7, 0x2b, 0x68, 0x26, 0x2e, 0x15, 0x87, 0x89, 0xfa, 0x10, 0xd5, 0x75, 0x44, 0x59, 0x1c, 0xa2,
	0xea, 0x6b, 0xd8, 0x57, 0x22, 0xe0, 0x46, 0x7e, 0xb6, 0x8f, 0x91, 0xfb, 0x81, 0x0c, 0x05, 0x6d,
	0xa1, 0xd0, 0xc9, 0x88, 0x61, 0x8a, 0x5b, 0x31, 0x9f, 0xeb, 0x28, 0x88, 0x4d, 0x4e, 0x7c, 0x94,
	0x88, 0x33, 0x62, 0x23, 0x26, 0xb0, 0x63, 0x5f, 0x2b, 0x3d, 0x46, 0x1e, 0xcf, 0xb6, 0x13, 0x32,
	0x5c, 0xc7, 0xc6, 0xd3, 0xf2, 0x8b, 0xa0, 0x34, 0xe9, 0x04, 0x22, 0x53, 0xf9, 0x45, 0x90, 0xaf,
	0x60, 0x4f, 0x86, 0xb7, 0x62, 0x61, 0x3c, 0x7c, 0xf4, 0x81, 0x9c, 0xd3, 0x27, 0x9d, 0x42, 0xaf,
	0xc2, 0x1a, 0x09, 0x6c, 0xe7, 0x6c, 0x2c, 0xe7, 0x64, 0x08, 0xce, 0x66, 0x59, 0x28, 0xf1, 0x5b,
	0x2c, 0xb4, 0xa1, 0xed, 0x4e, 0xe1, 0xc1, 0x54, 0x66, 0x23, 0xf5, 0x01, 0xd7, 0x09, 0xdb, 0xcb,
	0x2c, 0x2c, 0x71, 0x90, 0x36, 0x54, 0xd6, 0x4a, 0x46, 0x4a, 0x9a, 0x3b, 0xfa, 0xbf, 0x4e, 0xa1,
	0xd7, 0x60, 0x9b, 0x98, 0xfc, 0x04, 0x55, 0xbc, 0x82, 0xb9, 0x5b, 0x0b, 0xfa, 0x14, 0x97, 0x40,
	0xf7, 0xdf, 0x97, 0xc0, 0xec, 0x6e, 0x2d, 0x58, 0xc5, 0xa4, 0x27, 0xf2, 0x23, 0xd4, 0xf3, 0x3b,
	0x8b, 0x3e, 0xc3, 0xeb, 0xb5, 0x1f, 0xe4, 0xb8, 0xb0, 0x12, 0x86, 0x0a, 0x56, 0x5b, 0x6e, 0x03,
	0xd2, 0x81, 0xba, 0x51, 0x7c, 0x21, 0xbc, 0x74, 0x97, 0x3c, 0x4f, 0x86, 0x02, 0xb1, 0x9f, 0x71,
	0xa1, 0x74, 0xa1, 0x91, 0x28, 0xf0, 0x9e, 0xd2, 0xa7, 0x27, 0x28, 0xa9, 0x21, 0x68, 0x73, 0x8f,
	0x7c, 0xd2, 0x83, 0x32, 0x2e, 0x60, 0x4d, 0x3b, 0xf8, 0x66, 0x9d, 0x5c, 0xf9, 0xb1, 0x25, 0x58,
	0xca, 0x93, 0x3e, 0xec, 0xdd, 0x5f, 0xd5, 0x9a, 0xbe, 0x40, 0x0b, 0x7d, 0x68, 0x99, 0xa6, 0x02,
	0xd6, 0x0c, 0xf2, 0xa1, 0x26, 0xdf, 0x40, 0x25, 0xdb, 0xc8, 0xb4, 0x8b, 0xbf, 0xf6, 0x20, 0xe7,
	0xed, 0xa7, 0x14, 0xdb, 0x88, 0xc8, 0x53, 0xa8, 0xda, 0x65, 0xab, 0xd7, 0x7c, 0x21, 0xe8, 0x4b,
	0x7c, 0x59, 0x5b, 0x80, 0x7c, 0x07, 0xb0, 0x5d, 0xe8, 0xf4, 0x15, 0x5e, 0xa6, 0x95, 0x6f, 0xdf,
	0x86, 0x64, 0x39, 0x61, 0xf7, 0xcf, 0x02, 0x54, 0x37, 0x3b, 0x99, 0xd4, 0x60, 0xf7, 0x9c, 0xb9,
	0xfd, 0x99, 0x3b, 0x74, 0x1e, 0x91, 0x3a, 0x54, 0x06, 0xe3, 0xab, 0xf3, 0x77, 0xa3, 0xc9, 0xa5,
	0x53, 0xb0, 0x11, 0xbb, 0x9e, 0x4c, 0xfa, 0x83, 0xb1, 0xeb, 0x3c, 0xb6, 0x51, 0x7f, 0x3a, 0x1d,
	0x5d, 0x4e, 0xdc, 0xa1, 0x53, 0xb4, 0x36, 0xcb, 0x59, 0xe1, 0x0e, 0x69, 0x40, 0xf5, 0xfc, 0xea,
	0x97, 0xf7, 0x63, 0xd7, 0x66, 0x29, 0x11, 0x80, 0xf2, 0x45, 0x7f, 0x34, 0x76, 0x87, 0x4e, 0xd9,
	0xea, 0xfa, 0x83, 0x2b, 0x66, 0x89, 0x5d, 0xab, 0x1b, 0xba, 0x63, 0xf7, 0x12, 0xab, 0x55, 0x2c,
	0x77, 0x3d, 0x79, 0x37, 0xb9, 0xfa, 0x75, 0xe2, 0x54, 0xbb, 0x6f, 0xa0, 0x92, 0xcd, 0x08, 0xa9,
	0x42, 0x69, 0xfa, 0xd6, 0x75, 0xdf, 0x3b, 0x8f, 0x6c, 0x2e, 0xd6, 0x1f, 0x0c, 0x46, 0x33, 0xa7,
	0x60, 0xe1, 0xa1, 0xfb, 0x61, 0x34, 0x76, 0x1e, 0x5b, 0x78, 0x76, 0xcd, 0x66, 0x63, 0xd7, 0x29,
	0xce, 0xcb, 0xf8, 0x4f, 0xf6, 0xfa, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x94, 0x55, 0x22, 0x7a,
	0x73, 0x07, 0x00, 0x00,
}
