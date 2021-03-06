
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

type TComboExItem struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// CN: 动态转换一个已存在的对象实例。或者使用Obj.As().<目标对象>。
// EN: Dynamically convert an existing object instance. Or use Obj.As().<Target object>.
func AsComboExItem(obj interface{}) *TComboExItem {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TComboExItem{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
// Deprecated: use AsComboExItem.
func ComboExItemFromInst(inst uintptr) *TComboExItem {
    return AsComboExItem(inst)
}

// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
// Deprecated: use AsComboExItem.
func ComboExItemFromObj(obj IObject) *TComboExItem {
    return AsComboExItem(obj)
}

// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsComboExItem.
func ComboExItemFromUnsafePointer(ptr unsafe.Pointer) *TComboExItem {
    return AsComboExItem(ptr)
}

// -------------------------- Deprecated end --------------------------
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (c *TComboExItem) Instance() uintptr {
    return c.instance
}

// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (c *TComboExItem) UnsafeAddr() unsafe.Pointer {
    return c.ptr
}

// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (c *TComboExItem) IsValid() bool {
    return c.instance != 0
}

// CN: 检测当前对象是否继承自目标对象。
// EN: Checks whether the current object is inherited from the target object.
func (c *TComboExItem) Is() TIs {
    return TIs(c.instance)
}

// CN: 动态转换当前对象为目标对象。
// EN: Dynamically convert the current object to the target object.
//func (c *TComboExItem) As() TAs {
//    return TAs(c.instance)
//}

// CN: 获取类信息指针。
// EN: Get class information pointer.
func TComboExItemClass() TClass {
    return ComboExItem_StaticClassType()
}

// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (c *TComboExItem) Assign(Source IObject) {
    ComboExItem_Assign(c.instance, CheckPtr(Source))
}

// CN: 获取类名路径。
// EN: Get the class name path.
func (c *TComboExItem) GetNamePath() string {
    return ComboExItem_GetNamePath(c.instance)
}

// CN: 丢弃当前对象。
// EN: Discard the current object.
func (c *TComboExItem) DisposeOf() {
    ComboExItem_DisposeOf(c.instance)
}

// CN: 获取类的类型信息。
// EN: Get class type information.
func (c *TComboExItem) ClassType() TClass {
    return ComboExItem_ClassType(c.instance)
}

// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (c *TComboExItem) ClassName() string {
    return ComboExItem_ClassName(c.instance)
}

// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (c *TComboExItem) InstanceSize() int32 {
    return ComboExItem_InstanceSize(c.instance)
}

// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (c *TComboExItem) InheritsFrom(AClass TClass) bool {
    return ComboExItem_InheritsFrom(c.instance, AClass)
}

// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (c *TComboExItem) Equals(Obj IObject) bool {
    return ComboExItem_Equals(c.instance, CheckPtr(Obj))
}

// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (c *TComboExItem) GetHashCode() int32 {
    return ComboExItem_GetHashCode(c.instance)
}

// CN: 文本类信息。
// EN: Text information.
func (c *TComboExItem) ToString() string {
    return ComboExItem_ToString(c.instance)
}

func (c *TComboExItem) Indent() int32 {
    return ComboExItem_GetIndent(c.instance)
}

func (c *TComboExItem) SetIndent(value int32) {
    ComboExItem_SetIndent(c.instance, value)
}

func (c *TComboExItem) OverlayImageIndex() int32 {
    return ComboExItem_GetOverlayImageIndex(c.instance)
}

func (c *TComboExItem) SetOverlayImageIndex(value int32) {
    ComboExItem_SetOverlayImageIndex(c.instance, value)
}

func (c *TComboExItem) SelectedImageIndex() int32 {
    return ComboExItem_GetSelectedImageIndex(c.instance)
}

func (c *TComboExItem) SetSelectedImageIndex(value int32) {
    ComboExItem_SetSelectedImageIndex(c.instance, value)
}

func (c *TComboExItem) Data() unsafe.Pointer {
    return ComboExItem_GetData(c.instance)
}

func (c *TComboExItem) SetData(value unsafe.Pointer) {
    ComboExItem_SetData(c.instance, value)
}

// CN: 获取控件标题。
// EN: Get the control title.
func (c *TComboExItem) Caption() string {
    return ComboExItem_GetCaption(c.instance)
}

// CN: 设置控件标题。
// EN: Set the control title.
func (c *TComboExItem) SetCaption(value string) {
    ComboExItem_SetCaption(c.instance, value)
}

// CN: 获取图像在images中的索引。
// EN: .
func (c *TComboExItem) ImageIndex() int32 {
    return ComboExItem_GetImageIndex(c.instance)
}

// CN: 设置图像在images中的索引。
// EN: .
func (c *TComboExItem) SetImageIndex(value int32) {
    ComboExItem_SetImageIndex(c.instance, value)
}

func (c *TComboExItem) Collection() *TCollection {
    return AsCollection(ComboExItem_GetCollection(c.instance))
}

func (c *TComboExItem) SetCollection(value *TCollection) {
    ComboExItem_SetCollection(c.instance, CheckPtr(value))
}

func (c *TComboExItem) Index() int32 {
    return ComboExItem_GetIndex(c.instance)
}

func (c *TComboExItem) SetIndex(value int32) {
    ComboExItem_SetIndex(c.instance, value)
}

func (c *TComboExItem) DisplayName() string {
    return ComboExItem_GetDisplayName(c.instance)
}

func (c *TComboExItem) SetDisplayName(value string) {
    ComboExItem_SetDisplayName(c.instance, value)
}

