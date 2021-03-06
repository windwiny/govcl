
//----------------------------------------
// The code is automatically generated by the GenlibVcl tool.
// Copyright © ying32. All Rights Reserved.
// 
// Licensed under Apache License 2.0
//
//----------------------------------------


package vcl


import (
	. "github.com/ying32/govcl/vcl/api"
    . "github.com/ying32/govcl/vcl/types"
    "unsafe"
)

type TAction struct {
    IComponent
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// CN: 创建一个新的对象。
// EN: Create a new object.
func NewAction(owner IComponent) *TAction {
    a := new(TAction)
    a.instance = Action_Create(CheckPtr(owner))
    a.ptr = unsafe.Pointer(a.instance)
    return a
}

// CN: 动态转换一个已存在的对象实例。或者使用Obj.As().<目标对象>。
// EN: Dynamically convert an existing object instance. Or use Obj.As().<Target object>.
func AsAction(obj interface{}) *TAction {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TAction{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
// Deprecated: use AsAction.
func ActionFromInst(inst uintptr) *TAction {
    return AsAction(inst)
}

// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
// Deprecated: use AsAction.
func ActionFromObj(obj IObject) *TAction {
    return AsAction(obj)
}

// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsAction.
func ActionFromUnsafePointer(ptr unsafe.Pointer) *TAction {
    return AsAction(ptr)
}

// -------------------------- Deprecated end --------------------------
// CN: 释放对象。
// EN: Free object.
func (a *TAction) Free() {
    if a.instance != 0 {
        Action_Free(a.instance)
        a.instance, a.ptr = 0, nullptr
    }
}

// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (a *TAction) Instance() uintptr {
    return a.instance
}

// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (a *TAction) UnsafeAddr() unsafe.Pointer {
    return a.ptr
}

// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (a *TAction) IsValid() bool {
    return a.instance != 0
}

// CN: 检测当前对象是否继承自目标对象。
// EN: Checks whether the current object is inherited from the target object.
func (a *TAction) Is() TIs {
    return TIs(a.instance)
}

// CN: 动态转换当前对象为目标对象。
// EN: Dynamically convert the current object to the target object.
//func (a *TAction) As() TAs {
//    return TAs(a.instance)
//}

// CN: 获取类信息指针。
// EN: Get class information pointer.
func TActionClass() TClass {
    return Action_StaticClassType()
}

// CN: 执行。
// EN: .
func (a *TAction) Execute() bool {
    return Action_Execute(a.instance)
}

// CN: 控件更新。
// EN: Update.
func (a *TAction) Update() bool {
    return Action_Update(a.instance)
}

// CN: 是否有父容器。
// EN: Is there a parent container.
func (a *TAction) HasParent() bool {
    return Action_HasParent(a.instance)
}

// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (a *TAction) FindComponent(AName string) *TComponent {
    return AsComponent(Action_FindComponent(a.instance, AName))
}

// CN: 获取类名路径。
// EN: Get the class name path.
func (a *TAction) GetNamePath() string {
    return Action_GetNamePath(a.instance)
}

// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (a *TAction) Assign(Source IObject) {
    Action_Assign(a.instance, CheckPtr(Source))
}

// CN: 丢弃当前对象。
// EN: Discard the current object.
func (a *TAction) DisposeOf() {
    Action_DisposeOf(a.instance)
}

// CN: 获取类的类型信息。
// EN: Get class type information.
func (a *TAction) ClassType() TClass {
    return Action_ClassType(a.instance)
}

// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (a *TAction) ClassName() string {
    return Action_ClassName(a.instance)
}

// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (a *TAction) InstanceSize() int32 {
    return Action_InstanceSize(a.instance)
}

// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (a *TAction) InheritsFrom(AClass TClass) bool {
    return Action_InheritsFrom(a.instance, AClass)
}

// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (a *TAction) Equals(Obj IObject) bool {
    return Action_Equals(a.instance, CheckPtr(Obj))
}

// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (a *TAction) GetHashCode() int32 {
    return Action_GetHashCode(a.instance)
}

// CN: 文本类信息。
// EN: Text information.
func (a *TAction) ToString() string {
    return Action_ToString(a.instance)
}

func (a *TAction) AutoCheck() bool {
    return Action_GetAutoCheck(a.instance)
}

func (a *TAction) SetAutoCheck(value bool) {
    Action_SetAutoCheck(a.instance, value)
}

// CN: 获取控件标题。
// EN: Get the control title.
func (a *TAction) Caption() string {
    return Action_GetCaption(a.instance)
}

// CN: 设置控件标题。
// EN: Set the control title.
func (a *TAction) SetCaption(value string) {
    Action_SetCaption(a.instance, value)
}

// CN: 获取是否选中。
// EN: .
func (a *TAction) Checked() bool {
    return Action_GetChecked(a.instance)
}

// CN: 设置是否选中。
// EN: .
func (a *TAction) SetChecked(value bool) {
    Action_SetChecked(a.instance, value)
}

// CN: 获取控件启用。
// EN: Get the control enabled.
func (a *TAction) Enabled() bool {
    return Action_GetEnabled(a.instance)
}

// CN: 设置控件启用。
// EN: Set the control enabled.
func (a *TAction) SetEnabled(value bool) {
    Action_SetEnabled(a.instance, value)
}

// CN: 获取团组索引。
// EN: .
func (a *TAction) GroupIndex() int32 {
    return Action_GetGroupIndex(a.instance)
}

// CN: 设置团组索引。
// EN: .
func (a *TAction) SetGroupIndex(value int32) {
    Action_SetGroupIndex(a.instance, value)
}

// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (a *TAction) Hint() string {
    return Action_GetHint(a.instance)
}

// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (a *TAction) SetHint(value string) {
    Action_SetHint(a.instance, value)
}

// CN: 获取图像在images中的索引。
// EN: .
func (a *TAction) ImageIndex() int32 {
    return Action_GetImageIndex(a.instance)
}

// CN: 设置图像在images中的索引。
// EN: .
func (a *TAction) SetImageIndex(value int32) {
    Action_SetImageIndex(a.instance, value)
}

// CN: 获取快捷键。
// EN: .
func (a *TAction) ShortCut() TShortCut {
    return Action_GetShortCut(a.instance)
}

// CN: 设置快捷键。
// EN: .
func (a *TAction) SetShortCut(value TShortCut) {
    Action_SetShortCut(a.instance, value)
}

// CN: 获取控件可视。
// EN: Get the control visible.
func (a *TAction) Visible() bool {
    return Action_GetVisible(a.instance)
}

// CN: 设置控件可视。
// EN: Set the control visible.
func (a *TAction) SetVisible(value bool) {
    Action_SetVisible(a.instance, value)
}

// CN: 设置动作执行事件。
// EN: .
func (a *TAction) SetOnExecute(fn TNotifyEvent) {
    Action_SetOnExecute(a.instance, fn)
}

// CN: 设置动作更新事件。
// EN: .
func (a *TAction) SetOnUpdate(fn TNotifyEvent) {
    Action_SetOnUpdate(a.instance, fn)
}

// CN: 获取图标索引列表对象。
// EN: .
func (a *TAction) Images() *TImageList {
    return AsImageList(Action_GetImages(a.instance))
}

func (a *TAction) Index() int32 {
    return Action_GetIndex(a.instance)
}

func (a *TAction) SetIndex(value int32) {
    Action_SetIndex(a.instance, value)
}

// CN: 获取组件总数。
// EN: Get the total number of components.
func (a *TAction) ComponentCount() int32 {
    return Action_GetComponentCount(a.instance)
}

// CN: 获取组件索引。
// EN: Get component index.
func (a *TAction) ComponentIndex() int32 {
    return Action_GetComponentIndex(a.instance)
}

// CN: 设置组件索引。
// EN: Set component index.
func (a *TAction) SetComponentIndex(value int32) {
    Action_SetComponentIndex(a.instance, value)
}

// CN: 获取组件所有者。
// EN: Get component owner.
func (a *TAction) Owner() *TComponent {
    return AsComponent(Action_GetOwner(a.instance))
}

// CN: 获取组件名称。
// EN: Get the component name.
func (a *TAction) Name() string {
    return Action_GetName(a.instance)
}

// CN: 设置组件名称。
// EN: Set the component name.
func (a *TAction) SetName(value string) {
    Action_SetName(a.instance, value)
}

// CN: 获取对象标记。
// EN: Get the control tag.
func (a *TAction) Tag() int {
    return Action_GetTag(a.instance)
}

// CN: 设置对象标记。
// EN: Set the control tag.
func (a *TAction) SetTag(value int) {
    Action_SetTag(a.instance, value)
}

// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (a *TAction) Components(AIndex int32) *TComponent {
    return AsComponent(Action_GetComponents(a.instance, AIndex))
}

