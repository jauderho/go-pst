// go-pst is a library for reading Personal Storage Table (.pst) files (written in Go/Golang).
//
// Copyright (C) 2022  Marten Mooij
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

//go:generate msgp -tests=false

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: task.proto

package properties

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Contains an index identifying one of a set of pre-defined text strings to be associated with the flag.
	FlagString *int32 `protobuf:"varint,1,opt,name=flag_string,json=flagString,proto3,oneof" json:"flag_string,omitempty" msg:"2676483"` // @gotags: msg:"2676483"
	// Indicates whether a time-flagged Message object is complete.
	PercentComplete *float64 `protobuf:"fixed64,3,opt,name=percent_complete,json=percentComplete,proto3,oneof" json:"percent_complete,omitempty" msg:"2631705"` // @gotags: msg:"2631705"
	// Indicates the acceptance state of the task.
	TaskAcceptanceState *int32 `protobuf:"varint,4,opt,name=task_acceptance_state,json=taskAcceptanceState,proto3,oneof" json:"task_acceptance_state,omitempty" msg:"2632423"` // @gotags: msg:"2632423"
	// Indicates whether a task assignee has replied to a task request for this Task object.
	TaskAccepted *bool `protobuf:"varint,5,opt,name=task_accepted,json=taskAccepted,proto3,oneof" json:"task_accepted,omitempty" msg:"26317611"` // @gotags: msg:"26317611"
	// Indicates the number of minutes that the user actually spent working on a task.
	TaskActualEffort *int32 `protobuf:"varint,6,opt,name=task_actual_effort,json=taskActualEffort,proto3,oneof" json:"task_actual_effort,omitempty" msg:"2632003"` // @gotags: msg:"2632003"
	// Specifies the name of the user that last assigned the task.
	TaskAssigner *string `protobuf:"bytes,7,opt,name=task_assigner,json=taskAssigner,proto3,oneof" json:"task_assigner,omitempty" msg:"26323331"` // @gotags: msg:"26323331"
	// Indicates that the task is complete.
	TaskComplete *bool `protobuf:"varint,9,opt,name=task_complete,json=taskComplete,proto3,oneof" json:"task_complete,omitempty" msg:"26321211"` // @gotags: msg:"26321211"
	// The client can set this property, but it has no impact on the Task-Related Objects Protocol and is ignored by the server.
	TaskCustomFlags *int32 `protobuf:"varint,10,opt,name=task_custom_flags,json=taskCustomFlags,proto3,oneof" json:"task_custom_flags,omitempty" msg:"2632733"` // @gotags: msg:"2632733"
	// Specifies the date when the user completed work on the task.
	TaskDateCompleted *int64 `protobuf:"varint,11,opt,name=task_date_completed,json=taskDateCompleted,proto3,oneof" json:"task_date_completed,omitempty" msg:"26318364"` // @gotags: msg:"26318364"
	// Indicates whether new occurrences remain to be generated.
	TaskDeadOccurrence *bool `protobuf:"varint,12,opt,name=task_dead_occurrence,json=taskDeadOccurrence,proto3,oneof" json:"task_dead_occurrence,omitempty" msg:"26317711"` // @gotags: msg:"26317711"
	// Specifies the date by which the user expects work on the task to be complete.
	TaskDueDate *int64 `protobuf:"varint,13,opt,name=task_due_date,json=taskDueDate,proto3,oneof" json:"task_due_date,omitempty" msg:"26317364"` // @gotags: msg:"26317364"
	// Indicates the number of minutes that the user expects to work on a task.
	TaskEstimatedEffort *int32 `protobuf:"varint,14,opt,name=task_estimated_effort,json=taskEstimatedEffort,proto3,oneof" json:"task_estimated_effort,omitempty" msg:"2632013"` // @gotags: msg:"2632013"
	// Indicates that the Task object was originally created by the action of the current user or user agent instead of by the processing of a task request.
	TaskfCreator *bool `protobuf:"varint,15,opt,name=taskf_creator,json=taskfCreator,proto3,oneof" json:"taskf_creator,omitempty" msg:"26321411"` // @gotags: msg:"26321411"
	// Indicates the accuracy of the PidLidTaskOwner property (section 2.328).
	TaskfFixOffline *bool `protobuf:"varint,16,opt,name=taskf_fix_offline,json=taskfFixOffline,proto3,oneof" json:"taskf_fix_offline,omitempty" msg:"26324411"` // @gotags: msg:"26324411"
	// Indicates whether the task includes a recurrence pattern.
	TaskfRecurring *bool `protobuf:"varint,17,opt,name=taskf_recurring,json=taskfRecurring,proto3,oneof" json:"taskf_recurring,omitempty" msg:"26323811"` // @gotags: msg:"26323811"
	// Indicates the type of change that was last made to the Task object.
	TaskHistory *int32 `protobuf:"varint,19,opt,name=task_history,json=taskHistory,proto3,oneof" json:"task_history,omitempty" msg:"2632103"` // @gotags: msg:"2632103"
	// Contains the name of the user who most recently assigned the task, or the user to whom it was most recently assigned.
	TaskLastDelegate *string `protobuf:"bytes,20,opt,name=task_last_delegate,json=taskLastDelegate,proto3,oneof" json:"task_last_delegate,omitempty" msg:"26323731"` // @gotags: msg:"26323731"
	// Contains the date and time of the most recent change made to the Task object.
	TaskLastUpdate *int64 `protobuf:"varint,21,opt,name=task_last_update,json=taskLastUpdate,proto3,oneof" json:"task_last_update,omitempty" msg:"26320564"` // @gotags: msg:"26320564"
	// Contains the name of the most recent user to have been the owner of the task.
	TaskLastUser *string `protobuf:"bytes,22,opt,name=task_last_user,json=taskLastUser,proto3,oneof" json:"task_last_user,omitempty" msg:"26323431"` // @gotags: msg:"26323431"
	// Specifies the assignment status of the embedded Task object.
	TaskMode *int32 `protobuf:"varint,23,opt,name=task_mode,json=taskMode,proto3,oneof" json:"task_mode,omitempty" msg:"2673043"` // @gotags: msg:"2673043"
	// Provides optimization hints about the recipients of a Task object.
	TaskMultipleRecipients *int32 `protobuf:"varint,24,opt,name=task_multiple_recipients,json=taskMultipleRecipients,proto3,oneof" json:"task_multiple_recipients,omitempty" msg:"2632323"` // @gotags: msg:"2632323"
	// Not used. The client can set this property, but it has no impact on the Task-Related Objects Protocol and is ignored by the server.
	TaskNoCompute *bool `protobuf:"varint,25,opt,name=task_no_compute,json=taskNoCompute,proto3,oneof" json:"task_no_compute,omitempty" msg:"26323611"` // @gotags: msg:"26323611"
	// Provides an aid to custom sorting of Task objects.
	TaskOrdinal *int32 `protobuf:"varint,26,opt,name=task_ordinal,json=taskOrdinal,proto3,oneof" json:"task_ordinal,omitempty" msg:"2632353"` // @gotags: msg:"2632353"
	// Contains the name of the owner of the task.
	TaskOwner *string `protobuf:"bytes,27,opt,name=task_owner,json=taskOwner,proto3,oneof" json:"task_owner,omitempty" msg:"26321531"` // @gotags: msg:"26321531"
	// Indicates the role of the current user relative to the Task object.
	TaskOwnership *int32 `protobuf:"varint,28,opt,name=task_ownership,json=taskOwnership,proto3,oneof" json:"task_ownership,omitempty" msg:"2632413"` // @gotags: msg:"2632413"
	// Indicates whether future instances of recurring tasks need reminders, even though the value of the PidLidReminderSet property (section 2.222) is 0x00.
	TaskResetReminder *bool `protobuf:"varint,30,opt,name=task_reset_reminder,json=taskResetReminder,proto3,oneof" json:"task_reset_reminder,omitempty" msg:"26317511"` // @gotags: msg:"26317511"
	// Not used. The client can set this property, but it has no impact on the Task-Related Objects Protocol and is ignored by the server.
	TaskRole *string `protobuf:"bytes,31,opt,name=task_role,json=taskRole,proto3,oneof" json:"task_role,omitempty" msg:"26323931"` // @gotags: msg:"26323931"
	// Specifies the date on which the user expects work on the task to begin.
	TaskStartDate *int64 `protobuf:"varint,32,opt,name=task_start_date,json=taskStartDate,proto3,oneof" json:"task_start_date,omitempty" msg:"26317264"` // @gotags: msg:"26317264"
	// Indicates the current assignment state of the Task object.
	TaskState *int32 `protobuf:"varint,33,opt,name=task_state,json=taskState,proto3,oneof" json:"task_state,omitempty" msg:"2632033"` // @gotags: msg:"2632033"
	// Specifies the status of a task.
	TaskStatus *int32 `protobuf:"varint,34,opt,name=task_status,json=taskStatus,proto3,oneof" json:"task_status,omitempty" msg:"2631693"` // @gotags: msg:"2631693"
	// Indicates whether the task assignee has been requested to send an email message update upon completion of the assigned task.
	TaskStatusOnComplete *bool `protobuf:"varint,35,opt,name=task_status_on_complete,json=taskStatusOnComplete,proto3,oneof" json:"task_status_on_complete,omitempty" msg:"26320911"` // @gotags: msg:"26320911"
	// Indicates whether the task assignee has been requested to send a task update when the assigned Task object changes.
	TaskUpdates *bool `protobuf:"varint,36,opt,name=task_updates,json=taskUpdates,proto3,oneof" json:"task_updates,omitempty" msg:"26321111"` // @gotags: msg:"26321111"
	// Indicates which copy is the latest update of a Task object.
	TaskVersion *int32 `protobuf:"varint,37,opt,name=task_version,json=taskVersion,proto3,oneof" json:"task_version,omitempty" msg:"2632023"` // @gotags: msg:"2632023"
	// This property is set by the client but is ignored by the server.
	TeamTask *bool `protobuf:"varint,38,opt,name=team_task,json=teamTask,proto3,oneof" json:"team_task,omitempty" msg:"26317111"` // @gotags: msg:"26317111"
	// Contains the current time, in UTC, which is used to determine the sort order of objects in a consolidated to-do list.
	ToDoOrdinalDate *int64 `protobuf:"varint,39,opt,name=to_do_ordinal_date,json=toDoOrdinalDate,proto3,oneof" json:"to_do_ordinal_date,omitempty" msg:"26758464"` // @gotags: msg:"26758464"
	// Contains the numerals 0 through 9 that are used to break a tie when the PidLidToDoOrdinalDate property (section 2.344) is used to perform a sort of objects.
	ToDoSubOrdinal *string `protobuf:"bytes,40,opt,name=to_do_sub_ordinal,json=toDoSubOrdinal,proto3,oneof" json:"to_do_sub_ordinal,omitempty" msg:"26758531"` // @gotags: msg:"26758531"
	// Contains user-specifiable text to identify this Message object in a consolidated to-do list.
	ToDoTitle *string `protobuf:"bytes,41,opt,name=to_do_title,json=toDoTitle,proto3,oneof" json:"to_do_title,omitempty" msg:"26758831"` // @gotags: msg:"26758831"
	// Contains the value of the PidTagMessageDeliveryTime  property (section 2.789) when modifying the PidLidFlagRequest property (section 2.136).
	ValidFlagStringProof *int64 `protobuf:"varint,42,opt,name=valid_flag_string_proof,json=validFlagStringProof,proto3,oneof" json:"valid_flag_string_proof,omitempty" msg:"26763164"` // @gotags: msg:"26763164"
	// Contains a positive number whose negative is less than or equal to the value of the PidLidTaskOrdinal property (section 2.327) of all of the Task objects in the folder.
	OrdinalMost *int32 `protobuf:"varint,43,opt,name=ordinal_most,json=ordinalMost,proto3,oneof" json:"ordinal_most,omitempty" msg:"140503"` // @gotags: msg:"140503"
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{0}
}

func (x *Task) GetFlagString() int32 {
	if x != nil && x.FlagString != nil {
		return *x.FlagString
	}
	return 0
}

func (x *Task) GetPercentComplete() float64 {
	if x != nil && x.PercentComplete != nil {
		return *x.PercentComplete
	}
	return 0
}

func (x *Task) GetTaskAcceptanceState() int32 {
	if x != nil && x.TaskAcceptanceState != nil {
		return *x.TaskAcceptanceState
	}
	return 0
}

func (x *Task) GetTaskAccepted() bool {
	if x != nil && x.TaskAccepted != nil {
		return *x.TaskAccepted
	}
	return false
}

func (x *Task) GetTaskActualEffort() int32 {
	if x != nil && x.TaskActualEffort != nil {
		return *x.TaskActualEffort
	}
	return 0
}

func (x *Task) GetTaskAssigner() string {
	if x != nil && x.TaskAssigner != nil {
		return *x.TaskAssigner
	}
	return ""
}

func (x *Task) GetTaskComplete() bool {
	if x != nil && x.TaskComplete != nil {
		return *x.TaskComplete
	}
	return false
}

func (x *Task) GetTaskCustomFlags() int32 {
	if x != nil && x.TaskCustomFlags != nil {
		return *x.TaskCustomFlags
	}
	return 0
}

func (x *Task) GetTaskDateCompleted() int64 {
	if x != nil && x.TaskDateCompleted != nil {
		return *x.TaskDateCompleted
	}
	return 0
}

func (x *Task) GetTaskDeadOccurrence() bool {
	if x != nil && x.TaskDeadOccurrence != nil {
		return *x.TaskDeadOccurrence
	}
	return false
}

func (x *Task) GetTaskDueDate() int64 {
	if x != nil && x.TaskDueDate != nil {
		return *x.TaskDueDate
	}
	return 0
}

func (x *Task) GetTaskEstimatedEffort() int32 {
	if x != nil && x.TaskEstimatedEffort != nil {
		return *x.TaskEstimatedEffort
	}
	return 0
}

func (x *Task) GetTaskfCreator() bool {
	if x != nil && x.TaskfCreator != nil {
		return *x.TaskfCreator
	}
	return false
}

func (x *Task) GetTaskfFixOffline() bool {
	if x != nil && x.TaskfFixOffline != nil {
		return *x.TaskfFixOffline
	}
	return false
}

func (x *Task) GetTaskfRecurring() bool {
	if x != nil && x.TaskfRecurring != nil {
		return *x.TaskfRecurring
	}
	return false
}

func (x *Task) GetTaskHistory() int32 {
	if x != nil && x.TaskHistory != nil {
		return *x.TaskHistory
	}
	return 0
}

func (x *Task) GetTaskLastDelegate() string {
	if x != nil && x.TaskLastDelegate != nil {
		return *x.TaskLastDelegate
	}
	return ""
}

func (x *Task) GetTaskLastUpdate() int64 {
	if x != nil && x.TaskLastUpdate != nil {
		return *x.TaskLastUpdate
	}
	return 0
}

func (x *Task) GetTaskLastUser() string {
	if x != nil && x.TaskLastUser != nil {
		return *x.TaskLastUser
	}
	return ""
}

func (x *Task) GetTaskMode() int32 {
	if x != nil && x.TaskMode != nil {
		return *x.TaskMode
	}
	return 0
}

func (x *Task) GetTaskMultipleRecipients() int32 {
	if x != nil && x.TaskMultipleRecipients != nil {
		return *x.TaskMultipleRecipients
	}
	return 0
}

func (x *Task) GetTaskNoCompute() bool {
	if x != nil && x.TaskNoCompute != nil {
		return *x.TaskNoCompute
	}
	return false
}

func (x *Task) GetTaskOrdinal() int32 {
	if x != nil && x.TaskOrdinal != nil {
		return *x.TaskOrdinal
	}
	return 0
}

func (x *Task) GetTaskOwner() string {
	if x != nil && x.TaskOwner != nil {
		return *x.TaskOwner
	}
	return ""
}

func (x *Task) GetTaskOwnership() int32 {
	if x != nil && x.TaskOwnership != nil {
		return *x.TaskOwnership
	}
	return 0
}

func (x *Task) GetTaskResetReminder() bool {
	if x != nil && x.TaskResetReminder != nil {
		return *x.TaskResetReminder
	}
	return false
}

func (x *Task) GetTaskRole() string {
	if x != nil && x.TaskRole != nil {
		return *x.TaskRole
	}
	return ""
}

func (x *Task) GetTaskStartDate() int64 {
	if x != nil && x.TaskStartDate != nil {
		return *x.TaskStartDate
	}
	return 0
}

func (x *Task) GetTaskState() int32 {
	if x != nil && x.TaskState != nil {
		return *x.TaskState
	}
	return 0
}

func (x *Task) GetTaskStatus() int32 {
	if x != nil && x.TaskStatus != nil {
		return *x.TaskStatus
	}
	return 0
}

func (x *Task) GetTaskStatusOnComplete() bool {
	if x != nil && x.TaskStatusOnComplete != nil {
		return *x.TaskStatusOnComplete
	}
	return false
}

func (x *Task) GetTaskUpdates() bool {
	if x != nil && x.TaskUpdates != nil {
		return *x.TaskUpdates
	}
	return false
}

func (x *Task) GetTaskVersion() int32 {
	if x != nil && x.TaskVersion != nil {
		return *x.TaskVersion
	}
	return 0
}

func (x *Task) GetTeamTask() bool {
	if x != nil && x.TeamTask != nil {
		return *x.TeamTask
	}
	return false
}

func (x *Task) GetToDoOrdinalDate() int64 {
	if x != nil && x.ToDoOrdinalDate != nil {
		return *x.ToDoOrdinalDate
	}
	return 0
}

func (x *Task) GetToDoSubOrdinal() string {
	if x != nil && x.ToDoSubOrdinal != nil {
		return *x.ToDoSubOrdinal
	}
	return ""
}

func (x *Task) GetToDoTitle() string {
	if x != nil && x.ToDoTitle != nil {
		return *x.ToDoTitle
	}
	return ""
}

func (x *Task) GetValidFlagStringProof() int64 {
	if x != nil && x.ValidFlagStringProof != nil {
		return *x.ValidFlagStringProof
	}
	return 0
}

func (x *Task) GetOrdinalMost() int32 {
	if x != nil && x.OrdinalMost != nil {
		return *x.OrdinalMost
	}
	return 0
}

var File_task_proto protoreflect.FileDescriptor

var file_task_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfc, 0x13, 0x0a,
	0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x24, 0x0a, 0x0b, 0x66, 0x6c, 0x61, 0x67, 0x5f, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x0a, 0x66, 0x6c,
	0x61, 0x67, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x88, 0x01, 0x01, 0x12, 0x2e, 0x0a, 0x10, 0x70,
	0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x01, 0x48, 0x01, 0x52, 0x0f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x88, 0x01, 0x01, 0x12, 0x37, 0x0a, 0x15, 0x74,
	0x61, 0x73, 0x6b, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x48, 0x02, 0x52, 0x13, 0x74, 0x61,
	0x73, 0x6b, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x28, 0x0a, 0x0d, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x61, 0x63, 0x63,
	0x65, 0x70, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x48, 0x03, 0x52, 0x0c, 0x74,
	0x61, 0x73, 0x6b, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x88, 0x01, 0x01, 0x12, 0x31,
	0x0a, 0x12, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x61, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x65, 0x66,
	0x66, 0x6f, 0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x48, 0x04, 0x52, 0x10, 0x74, 0x61,
	0x73, 0x6b, 0x41, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x45, 0x66, 0x66, 0x6f, 0x72, 0x74, 0x88, 0x01,
	0x01, 0x12, 0x28, 0x0a, 0x0d, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e,
	0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x0c, 0x74, 0x61, 0x73, 0x6b,
	0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x28, 0x0a, 0x0d, 0x74,
	0x61, 0x73, 0x6b, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x08, 0x48, 0x06, 0x52, 0x0c, 0x74, 0x61, 0x73, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2f, 0x0a, 0x11, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05,
	0x48, 0x07, 0x52, 0x0f, 0x74, 0x61, 0x73, 0x6b, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x46, 0x6c,
	0x61, 0x67, 0x73, 0x88, 0x01, 0x01, 0x12, 0x33, 0x0a, 0x13, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x03, 0x48, 0x08, 0x52, 0x11, 0x74, 0x61, 0x73, 0x6b, 0x44, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x88, 0x01, 0x01, 0x12, 0x35, 0x0a, 0x14, 0x74,
	0x61, 0x73, 0x6b, 0x5f, 0x64, 0x65, 0x61, 0x64, 0x5f, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x48, 0x09, 0x52, 0x12, 0x74, 0x61, 0x73,
	0x6b, 0x44, 0x65, 0x61, 0x64, 0x4f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x88,
	0x01, 0x01, 0x12, 0x27, 0x0a, 0x0d, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x64, 0x75, 0x65, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x48, 0x0a, 0x52, 0x0b, 0x74, 0x61, 0x73,
	0x6b, 0x44, 0x75, 0x65, 0x44, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x12, 0x37, 0x0a, 0x15, 0x74,
	0x61, 0x73, 0x6b, 0x5f, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x65, 0x66,
	0x66, 0x6f, 0x72, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x48, 0x0b, 0x52, 0x13, 0x74, 0x61,
	0x73, 0x6b, 0x45, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x64, 0x45, 0x66, 0x66, 0x6f, 0x72,
	0x74, 0x88, 0x01, 0x01, 0x12, 0x28, 0x0a, 0x0d, 0x74, 0x61, 0x73, 0x6b, 0x66, 0x5f, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x48, 0x0c, 0x52, 0x0c, 0x74,
	0x61, 0x73, 0x6b, 0x66, 0x43, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x12, 0x2f,
	0x0a, 0x11, 0x74, 0x61, 0x73, 0x6b, 0x66, 0x5f, 0x66, 0x69, 0x78, 0x5f, 0x6f, 0x66, 0x66, 0x6c,
	0x69, 0x6e, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x08, 0x48, 0x0d, 0x52, 0x0f, 0x74, 0x61, 0x73,
	0x6b, 0x66, 0x46, 0x69, 0x78, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x2c, 0x0a, 0x0f, 0x74, 0x61, 0x73, 0x6b, 0x66, 0x5f, 0x72, 0x65, 0x63, 0x75, 0x72, 0x72, 0x69,
	0x6e, 0x67, 0x18, 0x11, 0x20, 0x01, 0x28, 0x08, 0x48, 0x0e, 0x52, 0x0e, 0x74, 0x61, 0x73, 0x6b,
	0x66, 0x52, 0x65, 0x63, 0x75, 0x72, 0x72, 0x69, 0x6e, 0x67, 0x88, 0x01, 0x01, 0x12, 0x26, 0x0a,
	0x0c, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x13, 0x20,
	0x01, 0x28, 0x05, 0x48, 0x0f, 0x52, 0x0b, 0x74, 0x61, 0x73, 0x6b, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x88, 0x01, 0x01, 0x12, 0x31, 0x0a, 0x12, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x6c, 0x61,
	0x73, 0x74, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x10, 0x52, 0x10, 0x74, 0x61, 0x73, 0x6b, 0x4c, 0x61, 0x73, 0x74, 0x44, 0x65, 0x6c,
	0x65, 0x67, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x10, 0x74, 0x61, 0x73, 0x6b,
	0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x15, 0x20, 0x01,
	0x28, 0x03, 0x48, 0x11, 0x52, 0x0e, 0x74, 0x61, 0x73, 0x6b, 0x4c, 0x61, 0x73, 0x74, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a, 0x0e, 0x74, 0x61, 0x73, 0x6b, 0x5f,
	0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x12, 0x52, 0x0c, 0x74, 0x61, 0x73, 0x6b, 0x4c, 0x61, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x88,
	0x01, 0x01, 0x12, 0x20, 0x0a, 0x09, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18,
	0x17, 0x20, 0x01, 0x28, 0x05, 0x48, 0x13, 0x52, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x4d, 0x6f, 0x64,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x3d, 0x0a, 0x18, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x6d, 0x75, 0x6c,
	0x74, 0x69, 0x70, 0x6c, 0x65, 0x5f, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x18, 0x20, 0x01, 0x28, 0x05, 0x48, 0x14, 0x52, 0x16, 0x74, 0x61, 0x73, 0x6b, 0x4d, 0x75,
	0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x73,
	0x88, 0x01, 0x01, 0x12, 0x2b, 0x0a, 0x0f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x6e, 0x6f, 0x5f, 0x63,
	0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x18, 0x19, 0x20, 0x01, 0x28, 0x08, 0x48, 0x15, 0x52, 0x0d,
	0x74, 0x61, 0x73, 0x6b, 0x4e, 0x6f, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x26, 0x0a, 0x0c, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x6c,
	0x18, 0x1a, 0x20, 0x01, 0x28, 0x05, 0x48, 0x16, 0x52, 0x0b, 0x74, 0x61, 0x73, 0x6b, 0x4f, 0x72,
	0x64, 0x69, 0x6e, 0x61, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x74, 0x61, 0x73, 0x6b,
	0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x09, 0x48, 0x17, 0x52, 0x09,
	0x74, 0x61, 0x73, 0x6b, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x2a, 0x0a, 0x0e,
	0x74, 0x61, 0x73, 0x6b, 0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x18, 0x1c,
	0x20, 0x01, 0x28, 0x05, 0x48, 0x18, 0x52, 0x0d, 0x74, 0x61, 0x73, 0x6b, 0x4f, 0x77, 0x6e, 0x65,
	0x72, 0x73, 0x68, 0x69, 0x70, 0x88, 0x01, 0x01, 0x12, 0x33, 0x0a, 0x13, 0x74, 0x61, 0x73, 0x6b,
	0x5f, 0x72, 0x65, 0x73, 0x65, 0x74, 0x5f, 0x72, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x18,
	0x1e, 0x20, 0x01, 0x28, 0x08, 0x48, 0x19, 0x52, 0x11, 0x74, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73,
	0x65, 0x74, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a,
	0x09, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x1a, 0x52, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x52, 0x6f, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x2b, 0x0a, 0x0f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x20, 0x20, 0x01, 0x28, 0x03, 0x48, 0x1b, 0x52, 0x0d, 0x74, 0x61, 0x73, 0x6b,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a,
	0x74, 0x61, 0x73, 0x6b, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x21, 0x20, 0x01, 0x28, 0x05,
	0x48, 0x1c, 0x52, 0x09, 0x74, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x24, 0x0a, 0x0b, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x22, 0x20, 0x01, 0x28, 0x05, 0x48, 0x1d, 0x52, 0x0a, 0x74, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x88, 0x01, 0x01, 0x12, 0x3a, 0x0a, 0x17, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x65, 0x18, 0x23, 0x20, 0x01, 0x28, 0x08, 0x48, 0x1e, 0x52, 0x14, 0x74, 0x61, 0x73, 0x6b, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x4f, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x88,
	0x01, 0x01, 0x12, 0x26, 0x0a, 0x0c, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x73, 0x18, 0x24, 0x20, 0x01, 0x28, 0x08, 0x48, 0x1f, 0x52, 0x0b, 0x74, 0x61, 0x73, 0x6b,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x88, 0x01, 0x01, 0x12, 0x26, 0x0a, 0x0c, 0x74, 0x61,
	0x73, 0x6b, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x25, 0x20, 0x01, 0x28, 0x05,
	0x48, 0x20, 0x52, 0x0b, 0x74, 0x61, 0x73, 0x6b, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x88,
	0x01, 0x01, 0x12, 0x20, 0x0a, 0x09, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x18,
	0x26, 0x20, 0x01, 0x28, 0x08, 0x48, 0x21, 0x52, 0x08, 0x74, 0x65, 0x61, 0x6d, 0x54, 0x61, 0x73,
	0x6b, 0x88, 0x01, 0x01, 0x12, 0x30, 0x0a, 0x12, 0x74, 0x6f, 0x5f, 0x64, 0x6f, 0x5f, 0x6f, 0x72,
	0x64, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x27, 0x20, 0x01, 0x28, 0x03,
	0x48, 0x22, 0x52, 0x0f, 0x74, 0x6f, 0x44, 0x6f, 0x4f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x6c, 0x44,
	0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2e, 0x0a, 0x11, 0x74, 0x6f, 0x5f, 0x64, 0x6f, 0x5f,
	0x73, 0x75, 0x62, 0x5f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x6c, 0x18, 0x28, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x23, 0x52, 0x0e, 0x74, 0x6f, 0x44, 0x6f, 0x53, 0x75, 0x62, 0x4f, 0x72, 0x64, 0x69,
	0x6e, 0x61, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0b, 0x74, 0x6f, 0x5f, 0x64, 0x6f, 0x5f,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x29, 0x20, 0x01, 0x28, 0x09, 0x48, 0x24, 0x52, 0x09, 0x74,
	0x6f, 0x44, 0x6f, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x3a, 0x0a, 0x17, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x5f, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x18, 0x2a, 0x20, 0x01, 0x28, 0x03, 0x48, 0x25, 0x52, 0x14,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x46, 0x6c, 0x61, 0x67, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x50,
	0x72, 0x6f, 0x6f, 0x66, 0x88, 0x01, 0x01, 0x12, 0x26, 0x0a, 0x0c, 0x6f, 0x72, 0x64, 0x69, 0x6e,
	0x61, 0x6c, 0x5f, 0x6d, 0x6f, 0x73, 0x74, 0x18, 0x2b, 0x20, 0x01, 0x28, 0x05, 0x48, 0x26, 0x52,
	0x0b, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x6c, 0x4d, 0x6f, 0x73, 0x74, 0x88, 0x01, 0x01, 0x42,
	0x0e, 0x0a, 0x0c, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x42,
	0x13, 0x0a, 0x11, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x65, 0x42, 0x18, 0x0a, 0x16, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x61, 0x63,
	0x63, 0x65, 0x70, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x42, 0x10,
	0x0a, 0x0e, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64,
	0x42, 0x15, 0x0a, 0x13, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x61, 0x63, 0x74, 0x75, 0x61, 0x6c,
	0x5f, 0x65, 0x66, 0x66, 0x6f, 0x72, 0x74, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x74, 0x61, 0x73, 0x6b,
	0x5f, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x74, 0x61,
	0x73, 0x6b, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x14, 0x0a, 0x12, 0x5f,
	0x74, 0x61, 0x73, 0x6b, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x66, 0x6c, 0x61, 0x67,
	0x73, 0x42, 0x16, 0x0a, 0x14, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x42, 0x17, 0x0a, 0x15, 0x5f, 0x74, 0x61,
	0x73, 0x6b, 0x5f, 0x64, 0x65, 0x61, 0x64, 0x5f, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x64, 0x75, 0x65, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x42, 0x18, 0x0a, 0x16, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x65, 0x73,
	0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x65, 0x66, 0x66, 0x6f, 0x72, 0x74, 0x42, 0x10,
	0x0a, 0x0e, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x66, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72,
	0x42, 0x14, 0x0a, 0x12, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x66, 0x5f, 0x66, 0x69, 0x78, 0x5f, 0x6f,
	0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x66,
	0x5f, 0x72, 0x65, 0x63, 0x75, 0x72, 0x72, 0x69, 0x6e, 0x67, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x74,
	0x61, 0x73, 0x6b, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x42, 0x15, 0x0a, 0x13, 0x5f,
	0x74, 0x61, 0x73, 0x6b, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x74, 0x61, 0x73, 0x6b,
	0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x74,
	0x61, 0x73, 0x6b, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x42, 0x1b, 0x0a, 0x19, 0x5f, 0x74, 0x61, 0x73,
	0x6b, 0x5f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x5f, 0x72, 0x65, 0x63, 0x69, 0x70,
	0x69, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x6e,
	0x6f, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x74, 0x61,
	0x73, 0x6b, 0x5f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x6c, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x74,
	0x61, 0x73, 0x6b, 0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x74, 0x61,
	0x73, 0x6b, 0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x42, 0x16, 0x0a, 0x14,
	0x5f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x72, 0x65, 0x73, 0x65, 0x74, 0x5f, 0x72, 0x65, 0x6d, 0x69,
	0x6e, 0x64, 0x65, 0x72, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x72, 0x6f,
	0x6c, 0x65, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x1a, 0x0a, 0x18, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x65, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x73, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x74, 0x61, 0x73,
	0x6b, 0x42, 0x15, 0x0a, 0x13, 0x5f, 0x74, 0x6f, 0x5f, 0x64, 0x6f, 0x5f, 0x6f, 0x72, 0x64, 0x69,
	0x6e, 0x61, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x42, 0x14, 0x0a, 0x12, 0x5f, 0x74, 0x6f, 0x5f,
	0x64, 0x6f, 0x5f, 0x73, 0x75, 0x62, 0x5f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x6c, 0x42, 0x0e,
	0x0a, 0x0c, 0x5f, 0x74, 0x6f, 0x5f, 0x64, 0x6f, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x42, 0x1a,
	0x0a, 0x18, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x5f, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x6f,
	0x72, 0x64, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x6d, 0x6f, 0x73, 0x74, 0x42, 0x28, 0x5a, 0x26, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x6f, 0x6f, 0x69, 0x6a, 0x74,
	0x65, 0x63, 0x68, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x73, 0x74, 0x3b, 0x70, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x69, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_task_proto_rawDescOnce sync.Once
	file_task_proto_rawDescData = file_task_proto_rawDesc
)

func file_task_proto_rawDescGZIP() []byte {
	file_task_proto_rawDescOnce.Do(func() {
		file_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_task_proto_rawDescData)
	})
	return file_task_proto_rawDescData
}

var file_task_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_task_proto_goTypes = []interface{}{
	(*Task)(nil), // 0: Task
}
var file_task_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_task_proto_init() }
func file_task_proto_init() {
	if File_task_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_task_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_task_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_task_proto_goTypes,
		DependencyIndexes: file_task_proto_depIdxs,
		MessageInfos:      file_task_proto_msgTypes,
	}.Build()
	File_task_proto = out.File
	file_task_proto_rawDesc = nil
	file_task_proto_goTypes = nil
	file_task_proto_depIdxs = nil
}
