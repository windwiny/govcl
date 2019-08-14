
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "github.com/ying32/govcl/vcl/api"
    . "github.com/ying32/govcl/vcl/types"
    "unsafe"
)

type TFrame struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewFrame
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewFrame(owner IComponent) *TFrame {
    f := new(TFrame)
    f.instance = Frame_Create(CheckPtr(owner))
    f.ptr = unsafe.Pointer(f.instance)
    return f
}

// FrameFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func FrameFromInst(inst uintptr) *TFrame {
    f := new(TFrame)
    f.instance = inst
    f.ptr = unsafe.Pointer(inst)
    return f
}

// FrameFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func FrameFromObj(obj IObject) *TFrame {
    f := new(TFrame)
    f.instance = CheckPtr(obj)
    f.ptr = unsafe.Pointer(f.instance)
    return f
}

// FrameFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func FrameFromUnsafePointer(ptr unsafe.Pointer) *TFrame {
    f := new(TFrame)
    f.instance = uintptr(ptr)
    f.ptr = ptr
    return f
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (f *TFrame) Free() {
    if f.instance != 0 {
        Frame_Free(f.instance)
        f.instance = 0
        f.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (f *TFrame) Instance() uintptr {
    return f.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (f *TFrame) UnsafeAddr() unsafe.Pointer {
    return f.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (f *TFrame) IsValid() bool {
    return f.instance != 0
}

// TFrameClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TFrameClass() TClass {
    return Frame_StaticClassType()
}

// DisableAutoRange
func (f *TFrame) DisableAutoRange() {
    Frame_DisableAutoRange(f.instance)
}

// EnableAutoRange
func (f *TFrame) EnableAutoRange() {
    Frame_EnableAutoRange(f.instance)
}

// ScrollInView
func (f *TFrame) ScrollInView(AControl IControl) {
    Frame_ScrollInView(f.instance, CheckPtr(AControl))
}

// CanFocus
// CN: 是否可以获得焦点。
// EN: .
func (f *TFrame) CanFocus() bool {
    return Frame_CanFocus(f.instance)
}

// ContainsControl
// CN: 返回是否包含指定控件。
// EN: it's contain a specified control.
func (f *TFrame) ContainsControl(Control IControl) bool {
    return Frame_ContainsControl(f.instance, CheckPtr(Control))
}

// ControlAtPos
// CN: 返回指定坐标及相关属性位置控件。
// EN: Returns the specified coordinate and the relevant attribute position control..
func (f *TFrame) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return ControlFromInst(Frame_ControlAtPos(f.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// DisableAlign
// CN: 禁用控件的对齐。
// EN: Disable control alignment.
func (f *TFrame) DisableAlign() {
    Frame_DisableAlign(f.instance)
}

// EnableAlign
// CN: 启用控件对齐。
// EN: Enabled control alignment.
func (f *TFrame) EnableAlign() {
    Frame_EnableAlign(f.instance)
}

// FindChildControl
// CN: 查找子控件。
// EN: Find sub controls.
func (f *TFrame) FindChildControl(ControlName string) *TControl {
    return ControlFromInst(Frame_FindChildControl(f.instance, ControlName))
}

// FlipChildren
func (f *TFrame) FlipChildren(AllLevels bool) {
    Frame_FlipChildren(f.instance, AllLevels)
}

// Focused
// CN: 返回是否获取焦点。
// EN: Return to get focus.
func (f *TFrame) Focused() bool {
    return Frame_Focused(f.instance)
}

// HandleAllocated
// CN: 句柄是否已经分配。
// EN: Is the handle already allocated.
func (f *TFrame) HandleAllocated() bool {
    return Frame_HandleAllocated(f.instance)
}

// InsertControl
// CN: 插入一个控件。
// EN: Insert a control.
func (f *TFrame) InsertControl(AControl IControl) {
    Frame_InsertControl(f.instance, CheckPtr(AControl))
}

// Invalidate
// CN: 要求重绘。
// EN: Redraw.
func (f *TFrame) Invalidate() {
    Frame_Invalidate(f.instance)
}

// PaintTo
// CN: 绘画至指定DC。
// EN: Painting to the specified DC.
func (f *TFrame) PaintTo(DC HDC, X int32, Y int32) {
    Frame_PaintTo(f.instance, DC , X , Y)
}

// RemoveControl
// CN: 移除一个控件。
// EN: Remove a control.
func (f *TFrame) RemoveControl(AControl IControl) {
    Frame_RemoveControl(f.instance, CheckPtr(AControl))
}

// Realign
// CN: 重新对齐。
// EN: Realign.
func (f *TFrame) Realign() {
    Frame_Realign(f.instance)
}

// Repaint
// CN: 重绘。
// EN: Repaint.
func (f *TFrame) Repaint() {
    Frame_Repaint(f.instance)
}

// ScaleBy
// CN: 按比例缩放。
// EN: Scale by.
func (f *TFrame) ScaleBy(M int32, D int32) {
    Frame_ScaleBy(f.instance, M , D)
}

// ScrollBy
// CN: 滚动至指定位置。
// EN: Scroll by.
func (f *TFrame) ScrollBy(DeltaX int32, DeltaY int32) {
    Frame_ScrollBy(f.instance, DeltaX , DeltaY)
}

// SetBounds
// CN: 设置组件边界。
// EN: Set component boundaries.
func (f *TFrame) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    Frame_SetBounds(f.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
// CN: 设置控件焦点。
// EN: Set control focus.
func (f *TFrame) SetFocus() {
    Frame_SetFocus(f.instance)
}

// Update
// CN: 控件更新。
// EN: Update.
func (f *TFrame) Update() {
    Frame_Update(f.instance)
}

// UpdateControlState
// CN: 更新控件状态。
// EN: Update control status.
func (f *TFrame) UpdateControlState() {
    Frame_UpdateControlState(f.instance)
}

// BringToFront
// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (f *TFrame) BringToFront() {
    Frame_BringToFront(f.instance)
}

// ClientToScreen
// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (f *TFrame) ClientToScreen(Point TPoint) TPoint {
    return Frame_ClientToScreen(f.instance, Point)
}

// ClientToParent
// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (f *TFrame) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return Frame_ClientToParent(f.instance, Point , CheckPtr(AParent))
}

// Dragging
// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (f *TFrame) Dragging() bool {
    return Frame_Dragging(f.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (f *TFrame) HasParent() bool {
    return Frame_HasParent(f.instance)
}

// Hide
// CN: 隐藏控件。
// EN: Hidden control.
func (f *TFrame) Hide() {
    Frame_Hide(f.instance)
}

// Perform
// CN: 发送一个消息。
// EN: Send a message.
func (f *TFrame) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return Frame_Perform(f.instance, Msg , WParam , LParam)
}

// Refresh
// CN: 刷新控件。
// EN: Refresh control.
func (f *TFrame) Refresh() {
    Frame_Refresh(f.instance)
}

// ScreenToClient
// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (f *TFrame) ScreenToClient(Point TPoint) TPoint {
    return Frame_ScreenToClient(f.instance, Point)
}

// ParentToClient
// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (f *TFrame) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return Frame_ParentToClient(f.instance, Point , CheckPtr(AParent))
}

// SendToBack
// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (f *TFrame) SendToBack() {
    Frame_SendToBack(f.instance)
}

// Show
// CN: 显示控件。
// EN: Show control.
func (f *TFrame) Show() {
    Frame_Show(f.instance)
}

// GetTextBuf
// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (f *TFrame) GetTextBuf(Buffer *string, BufSize int32) int32 {
    return Frame_GetTextBuf(f.instance, Buffer , BufSize)
}

// GetTextLen
// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (f *TFrame) GetTextLen() int32 {
    return Frame_GetTextLen(f.instance)
}

// SetTextBuf
// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (f *TFrame) SetTextBuf(Buffer string) {
    Frame_SetTextBuf(f.instance, Buffer)
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (f *TFrame) FindComponent(AName string) *TComponent {
    return ComponentFromInst(Frame_FindComponent(f.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (f *TFrame) GetNamePath() string {
    return Frame_GetNamePath(f.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (f *TFrame) Assign(Source IObject) {
    Frame_Assign(f.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (f *TFrame) DisposeOf() {
    Frame_DisposeOf(f.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (f *TFrame) ClassType() TClass {
    return Frame_ClassType(f.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (f *TFrame) ClassName() string {
    return Frame_ClassName(f.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (f *TFrame) InstanceSize() int32 {
    return Frame_InstanceSize(f.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (f *TFrame) InheritsFrom(AClass TClass) bool {
    return Frame_InheritsFrom(f.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (f *TFrame) Equals(Obj IObject) bool {
    return Frame_Equals(f.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (f *TFrame) GetHashCode() int32 {
    return Frame_GetHashCode(f.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (f *TFrame) ToString() string {
    return Frame_ToString(f.instance)
}

// Align
// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (f *TFrame) Align() TAlign {
    return Frame_GetAlign(f.instance)
}

// SetAlign
// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (f *TFrame) SetAlign(value TAlign) {
    Frame_SetAlign(f.instance, value)
}

// Anchors
// CN: 获取四个角位置的锚点。
// EN: .
func (f *TFrame) Anchors() TAnchors {
    return Frame_GetAnchors(f.instance)
}

// SetAnchors
// CN: 设置四个角位置的锚点。
// EN: .
func (f *TFrame) SetAnchors(value TAnchors) {
    Frame_SetAnchors(f.instance, value)
}

// AutoScroll
func (f *TFrame) AutoScroll() bool {
    return Frame_GetAutoScroll(f.instance)
}

// SetAutoScroll
func (f *TFrame) SetAutoScroll(value bool) {
    Frame_SetAutoScroll(f.instance, value)
}

// AutoSize
// CN: 获取自动调整大小。
// EN: .
func (f *TFrame) AutoSize() bool {
    return Frame_GetAutoSize(f.instance)
}

// SetAutoSize
// CN: 设置自动调整大小。
// EN: .
func (f *TFrame) SetAutoSize(value bool) {
    Frame_SetAutoSize(f.instance, value)
}

// BiDiMode
func (f *TFrame) BiDiMode() TBiDiMode {
    return Frame_GetBiDiMode(f.instance)
}

// SetBiDiMode
func (f *TFrame) SetBiDiMode(value TBiDiMode) {
    Frame_SetBiDiMode(f.instance, value)
}

// Constraints
func (f *TFrame) Constraints() *TSizeConstraints {
    return SizeConstraintsFromInst(Frame_GetConstraints(f.instance))
}

// SetConstraints
func (f *TFrame) SetConstraints(value *TSizeConstraints) {
    Frame_SetConstraints(f.instance, CheckPtr(value))
}

// DockSite
// CN: 获取停靠站点。
// EN: Get Docking site.
func (f *TFrame) DockSite() bool {
    return Frame_GetDockSite(f.instance)
}

// SetDockSite
// CN: 设置停靠站点。
// EN: Set Docking site.
func (f *TFrame) SetDockSite(value bool) {
    Frame_SetDockSite(f.instance, value)
}

// DoubleBuffered
// CN: 获取设置控件双缓冲。
// EN: Get Set control double buffering.
func (f *TFrame) DoubleBuffered() bool {
    return Frame_GetDoubleBuffered(f.instance)
}

// SetDoubleBuffered
// CN: 设置设置控件双缓冲。
// EN: Set Set control double buffering.
func (f *TFrame) SetDoubleBuffered(value bool) {
    Frame_SetDoubleBuffered(f.instance, value)
}

// DragCursor
// CN: 获取设置控件拖拽时的光标。
// EN: Get Set the cursor when the control is dragged.
func (f *TFrame) DragCursor() TCursor {
    return Frame_GetDragCursor(f.instance)
}

// SetDragCursor
// CN: 设置设置控件拖拽时的光标。
// EN: Set Set the cursor when the control is dragged.
func (f *TFrame) SetDragCursor(value TCursor) {
    Frame_SetDragCursor(f.instance, value)
}

// DragKind
// CN: 获取拖拽方式。
// EN: Get Drag and drop.
func (f *TFrame) DragKind() TDragKind {
    return Frame_GetDragKind(f.instance)
}

// SetDragKind
// CN: 设置拖拽方式。
// EN: Set Drag and drop.
func (f *TFrame) SetDragKind(value TDragKind) {
    Frame_SetDragKind(f.instance, value)
}

// DragMode
// CN: 获取拖拽模式。
// EN: Get Drag mode.
func (f *TFrame) DragMode() TDragMode {
    return Frame_GetDragMode(f.instance)
}

// SetDragMode
// CN: 设置拖拽模式。
// EN: Set Drag mode.
func (f *TFrame) SetDragMode(value TDragMode) {
    Frame_SetDragMode(f.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (f *TFrame) Enabled() bool {
    return Frame_GetEnabled(f.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (f *TFrame) SetEnabled(value bool) {
    Frame_SetEnabled(f.instance, value)
}

// Color
// CN: 获取颜色。
// EN: Get color.
func (f *TFrame) Color() TColor {
    return Frame_GetColor(f.instance)
}

// SetColor
// CN: 设置颜色。
// EN: Set color.
func (f *TFrame) SetColor(value TColor) {
    Frame_SetColor(f.instance, value)
}

// Ctl3D
func (f *TFrame) Ctl3D() bool {
    return Frame_GetCtl3D(f.instance)
}

// SetCtl3D
func (f *TFrame) SetCtl3D(value bool) {
    Frame_SetCtl3D(f.instance, value)
}

// Font
// CN: 获取字体。
// EN: Get Font.
func (f *TFrame) Font() *TFont {
    return FontFromInst(Frame_GetFont(f.instance))
}

// SetFont
// CN: 设置字体。
// EN: Set Font.
func (f *TFrame) SetFont(value *TFont) {
    Frame_SetFont(f.instance, CheckPtr(value))
}

// ParentBackground
func (f *TFrame) ParentBackground() bool {
    return Frame_GetParentBackground(f.instance)
}

// SetParentBackground
func (f *TFrame) SetParentBackground(value bool) {
    Frame_SetParentBackground(f.instance, value)
}

// ParentColor
// CN: 获取父容器颜色。
// EN: Get parent color.
func (f *TFrame) ParentColor() bool {
    return Frame_GetParentColor(f.instance)
}

// SetParentColor
// CN: 设置父容器颜色。
// EN: Set parent color.
func (f *TFrame) SetParentColor(value bool) {
    Frame_SetParentColor(f.instance, value)
}

// ParentCtl3D
func (f *TFrame) ParentCtl3D() bool {
    return Frame_GetParentCtl3D(f.instance)
}

// SetParentCtl3D
func (f *TFrame) SetParentCtl3D(value bool) {
    Frame_SetParentCtl3D(f.instance, value)
}

// ParentDoubleBuffered
// CN: 获取父容器双缓冲。
// EN: Get Parent container double buffering.
func (f *TFrame) ParentDoubleBuffered() bool {
    return Frame_GetParentDoubleBuffered(f.instance)
}

// SetParentDoubleBuffered
// CN: 设置父容器双缓冲。
// EN: Set Parent container double buffering.
func (f *TFrame) SetParentDoubleBuffered(value bool) {
    Frame_SetParentDoubleBuffered(f.instance, value)
}

// ParentFont
// CN: 获取父容器字体。
// EN: Get Parent container font.
func (f *TFrame) ParentFont() bool {
    return Frame_GetParentFont(f.instance)
}

// SetParentFont
// CN: 设置父容器字体。
// EN: Set Parent container font.
func (f *TFrame) SetParentFont(value bool) {
    Frame_SetParentFont(f.instance, value)
}

// ParentShowHint
func (f *TFrame) ParentShowHint() bool {
    return Frame_GetParentShowHint(f.instance)
}

// SetParentShowHint
func (f *TFrame) SetParentShowHint(value bool) {
    Frame_SetParentShowHint(f.instance, value)
}

// PopupMenu
// CN: 获取右键菜单。
// EN: Get Right click menu.
func (f *TFrame) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(Frame_GetPopupMenu(f.instance))
}

// SetPopupMenu
// CN: 设置右键菜单。
// EN: Set Right click menu.
func (f *TFrame) SetPopupMenu(value IComponent) {
    Frame_SetPopupMenu(f.instance, CheckPtr(value))
}

// ShowHint
// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (f *TFrame) ShowHint() bool {
    return Frame_GetShowHint(f.instance)
}

// SetShowHint
// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (f *TFrame) SetShowHint(value bool) {
    Frame_SetShowHint(f.instance, value)
}

// TabOrder
// CN: 获取Tab切换顺序序号。
// EN: Get Tab switching sequence number.
func (f *TFrame) TabOrder() TTabOrder {
    return Frame_GetTabOrder(f.instance)
}

// SetTabOrder
// CN: 设置Tab切换顺序序号。
// EN: Set Tab switching sequence number.
func (f *TFrame) SetTabOrder(value TTabOrder) {
    Frame_SetTabOrder(f.instance, value)
}

// TabStop
// CN: 获取Tab可停留。
// EN: Get Tab can stay.
func (f *TFrame) TabStop() bool {
    return Frame_GetTabStop(f.instance)
}

// SetTabStop
// CN: 设置Tab可停留。
// EN: Set Tab can stay.
func (f *TFrame) SetTabStop(value bool) {
    Frame_SetTabStop(f.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (f *TFrame) Visible() bool {
    return Frame_GetVisible(f.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (f *TFrame) SetVisible(value bool) {
    Frame_SetVisible(f.instance, value)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (f *TFrame) SetOnClick(fn TNotifyEvent) {
    Frame_SetOnClick(f.instance, fn)
}

// SetOnContextPopup
// CN: 设置上下文弹出事件，一般是右键时弹出。
// EN: Set Context popup event, usually pop up when right click.
func (f *TFrame) SetOnContextPopup(fn TContextPopupEvent) {
    Frame_SetOnContextPopup(f.instance, fn)
}

// SetOnDblClick
// CN: 设置双击事件。
// EN: .
func (f *TFrame) SetOnDblClick(fn TNotifyEvent) {
    Frame_SetOnDblClick(f.instance, fn)
}

// SetOnDockDrop
func (f *TFrame) SetOnDockDrop(fn TDockDropEvent) {
    Frame_SetOnDockDrop(f.instance, fn)
}

// SetOnDragDrop
// CN: 设置拖拽下落事件。
// EN: Set Drag and drop event.
func (f *TFrame) SetOnDragDrop(fn TDragDropEvent) {
    Frame_SetOnDragDrop(f.instance, fn)
}

// SetOnDragOver
// CN: 设置拖拽完成事件。
// EN: Set Drag and drop completion event.
func (f *TFrame) SetOnDragOver(fn TDragOverEvent) {
    Frame_SetOnDragOver(f.instance, fn)
}

// SetOnEndDock
// CN: 设置停靠结束事件。
// EN: Set Dock end event.
func (f *TFrame) SetOnEndDock(fn TEndDragEvent) {
    Frame_SetOnEndDock(f.instance, fn)
}

// SetOnEndDrag
// CN: 设置拖拽结束。
// EN: Set End of drag.
func (f *TFrame) SetOnEndDrag(fn TEndDragEvent) {
    Frame_SetOnEndDrag(f.instance, fn)
}

// SetOnEnter
// CN: 设置焦点进入。
// EN: Set Focus entry.
func (f *TFrame) SetOnEnter(fn TNotifyEvent) {
    Frame_SetOnEnter(f.instance, fn)
}

// SetOnExit
// CN: 设置焦点退出。
// EN: Set Focus exit.
func (f *TFrame) SetOnExit(fn TNotifyEvent) {
    Frame_SetOnExit(f.instance, fn)
}

// SetOnGesture
func (f *TFrame) SetOnGesture(fn TGestureEvent) {
    Frame_SetOnGesture(f.instance, fn)
}

// SetOnGetSiteInfo
func (f *TFrame) SetOnGetSiteInfo(fn TGetSiteInfoEvent) {
    Frame_SetOnGetSiteInfo(f.instance, fn)
}

// SetOnMouseActivate
func (f *TFrame) SetOnMouseActivate(fn TMouseActivateEvent) {
    Frame_SetOnMouseActivate(f.instance, fn)
}

// SetOnMouseDown
// CN: 设置鼠标按下事件。
// EN: Set Mouse down event.
func (f *TFrame) SetOnMouseDown(fn TMouseEvent) {
    Frame_SetOnMouseDown(f.instance, fn)
}

// SetOnMouseEnter
// CN: 设置鼠标进入事件。
// EN: Set Mouse entry event.
func (f *TFrame) SetOnMouseEnter(fn TNotifyEvent) {
    Frame_SetOnMouseEnter(f.instance, fn)
}

// SetOnMouseLeave
// CN: 设置鼠标离开事件。
// EN: Set Mouse leave event.
func (f *TFrame) SetOnMouseLeave(fn TNotifyEvent) {
    Frame_SetOnMouseLeave(f.instance, fn)
}

// SetOnMouseMove
// CN: 设置鼠标移动事件。
// EN: .
func (f *TFrame) SetOnMouseMove(fn TMouseMoveEvent) {
    Frame_SetOnMouseMove(f.instance, fn)
}

// SetOnMouseUp
// CN: 设置鼠标抬起事件。
// EN: Set Mouse lift event.
func (f *TFrame) SetOnMouseUp(fn TMouseEvent) {
    Frame_SetOnMouseUp(f.instance, fn)
}

// SetOnMouseWheel
// CN: 设置鼠标滚轮事件。
// EN: .
func (f *TFrame) SetOnMouseWheel(fn TMouseWheelEvent) {
    Frame_SetOnMouseWheel(f.instance, fn)
}

// SetOnMouseWheelDown
// CN: 设置鼠标滚轮按下事件。
// EN: .
func (f *TFrame) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
    Frame_SetOnMouseWheelDown(f.instance, fn)
}

// SetOnMouseWheelUp
// CN: 设置鼠标滚轮抬起事件。
// EN: .
func (f *TFrame) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
    Frame_SetOnMouseWheelUp(f.instance, fn)
}

// SetOnResize
// CN: 设置大小被改变事件。
// EN: .
func (f *TFrame) SetOnResize(fn TNotifyEvent) {
    Frame_SetOnResize(f.instance, fn)
}

// SetOnStartDock
// CN: 设置启动停靠。
// EN: .
func (f *TFrame) SetOnStartDock(fn TStartDockEvent) {
    Frame_SetOnStartDock(f.instance, fn)
}

// SetOnUnDock
func (f *TFrame) SetOnUnDock(fn TUnDockEvent) {
    Frame_SetOnUnDock(f.instance, fn)
}

// HorzScrollBar
func (f *TFrame) HorzScrollBar() *TControlScrollBar {
    return ControlScrollBarFromInst(Frame_GetHorzScrollBar(f.instance))
}

// SetHorzScrollBar
func (f *TFrame) SetHorzScrollBar(value *TControlScrollBar) {
    Frame_SetHorzScrollBar(f.instance, CheckPtr(value))
}

// VertScrollBar
func (f *TFrame) VertScrollBar() *TControlScrollBar {
    return ControlScrollBarFromInst(Frame_GetVertScrollBar(f.instance))
}

// SetVertScrollBar
func (f *TFrame) SetVertScrollBar(value *TControlScrollBar) {
    Frame_SetVertScrollBar(f.instance, CheckPtr(value))
}

// DockClientCount
// CN: 获取依靠客户端总数。
// EN: .
func (f *TFrame) DockClientCount() int32 {
    return Frame_GetDockClientCount(f.instance)
}

// AlignDisabled
// CN: 获取禁用对齐。
// EN: .
func (f *TFrame) AlignDisabled() bool {
    return Frame_GetAlignDisabled(f.instance)
}

// MouseInClient
// CN: 获取鼠标是否在客户端，仅VCL有效。
// EN: Get Whether the mouse is on the client, only VCL is valid.
func (f *TFrame) MouseInClient() bool {
    return Frame_GetMouseInClient(f.instance)
}

// VisibleDockClientCount
// CN: 获取当前停靠的可视总数。
// EN: Get The total number of visible calls currently docked.
func (f *TFrame) VisibleDockClientCount() int32 {
    return Frame_GetVisibleDockClientCount(f.instance)
}

// Brush
// CN: 获取画刷对象。
// EN: Get Brush.
func (f *TFrame) Brush() *TBrush {
    return BrushFromInst(Frame_GetBrush(f.instance))
}

// ControlCount
// CN: 获取子控件数。
// EN: Get Number of child controls.
func (f *TFrame) ControlCount() int32 {
    return Frame_GetControlCount(f.instance)
}

// Handle
// CN: 获取控件句柄。
// EN: Get Control handle.
func (f *TFrame) Handle() HWND {
    return Frame_GetHandle(f.instance)
}

// ParentWindow
// CN: 获取父容器句柄。
// EN: Get Parent container handle.
func (f *TFrame) ParentWindow() HWND {
    return Frame_GetParentWindow(f.instance)
}

// SetParentWindow
// CN: 设置父容器句柄。
// EN: Set Parent container handle.
func (f *TFrame) SetParentWindow(value HWND) {
    Frame_SetParentWindow(f.instance, value)
}

// UseDockManager
// CN: 获取使用停靠管理。
// EN: .
func (f *TFrame) UseDockManager() bool {
    return Frame_GetUseDockManager(f.instance)
}

// SetUseDockManager
// CN: 设置使用停靠管理。
// EN: .
func (f *TFrame) SetUseDockManager(value bool) {
    Frame_SetUseDockManager(f.instance, value)
}

// Action
func (f *TFrame) Action() *TAction {
    return ActionFromInst(Frame_GetAction(f.instance))
}

// SetAction
func (f *TFrame) SetAction(value IComponent) {
    Frame_SetAction(f.instance, CheckPtr(value))
}

// BoundsRect
func (f *TFrame) BoundsRect() TRect {
    return Frame_GetBoundsRect(f.instance)
}

// SetBoundsRect
func (f *TFrame) SetBoundsRect(value TRect) {
    Frame_SetBoundsRect(f.instance, value)
}

// ClientHeight
// CN: 获取客户区高度。
// EN: Get client height.
func (f *TFrame) ClientHeight() int32 {
    return Frame_GetClientHeight(f.instance)
}

// SetClientHeight
// CN: 设置客户区高度。
// EN: Set client height.
func (f *TFrame) SetClientHeight(value int32) {
    Frame_SetClientHeight(f.instance, value)
}

// ClientOrigin
func (f *TFrame) ClientOrigin() TPoint {
    return Frame_GetClientOrigin(f.instance)
}

// ClientRect
// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (f *TFrame) ClientRect() TRect {
    return Frame_GetClientRect(f.instance)
}

// ClientWidth
// CN: 获取客户区宽度。
// EN: Get client width.
func (f *TFrame) ClientWidth() int32 {
    return Frame_GetClientWidth(f.instance)
}

// SetClientWidth
// CN: 设置客户区宽度。
// EN: Set client width.
func (f *TFrame) SetClientWidth(value int32) {
    Frame_SetClientWidth(f.instance, value)
}

// ControlState
// CN: 获取控件状态。
// EN: Get control state.
func (f *TFrame) ControlState() TControlState {
    return Frame_GetControlState(f.instance)
}

// SetControlState
// CN: 设置控件状态。
// EN: Set control state.
func (f *TFrame) SetControlState(value TControlState) {
    Frame_SetControlState(f.instance, value)
}

// ControlStyle
// CN: 获取控件样式。
// EN: Get control style.
func (f *TFrame) ControlStyle() TControlStyle {
    return Frame_GetControlStyle(f.instance)
}

// SetControlStyle
// CN: 设置控件样式。
// EN: Set control style.
func (f *TFrame) SetControlStyle(value TControlStyle) {
    Frame_SetControlStyle(f.instance, value)
}

// ExplicitLeft
func (f *TFrame) ExplicitLeft() int32 {
    return Frame_GetExplicitLeft(f.instance)
}

// ExplicitTop
func (f *TFrame) ExplicitTop() int32 {
    return Frame_GetExplicitTop(f.instance)
}

// ExplicitWidth
func (f *TFrame) ExplicitWidth() int32 {
    return Frame_GetExplicitWidth(f.instance)
}

// ExplicitHeight
func (f *TFrame) ExplicitHeight() int32 {
    return Frame_GetExplicitHeight(f.instance)
}

// Floating
func (f *TFrame) Floating() bool {
    return Frame_GetFloating(f.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (f *TFrame) Parent() *TWinControl {
    return WinControlFromInst(Frame_GetParent(f.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (f *TFrame) SetParent(value IWinControl) {
    Frame_SetParent(f.instance, CheckPtr(value))
}

// StyleElements
// CN: 获取样式元素。
// EN: Get Style element.
func (f *TFrame) StyleElements() TStyleElements {
    return Frame_GetStyleElements(f.instance)
}

// SetStyleElements
// CN: 设置样式元素。
// EN: Set Style element.
func (f *TFrame) SetStyleElements(value TStyleElements) {
    Frame_SetStyleElements(f.instance, value)
}

// AlignWithMargins
// CN: 获取边距，仅VCL有效。
// EN: Get The margin. Only VCL is valid..
func (f *TFrame) AlignWithMargins() bool {
    return Frame_GetAlignWithMargins(f.instance)
}

// SetAlignWithMargins
// CN: 设置边距，仅VCL有效。
// EN: Set The margin. Only VCL is valid..
func (f *TFrame) SetAlignWithMargins(value bool) {
    Frame_SetAlignWithMargins(f.instance, value)
}

// Left
// CN: 获取左边位置。
// EN: Get Left position.
func (f *TFrame) Left() int32 {
    return Frame_GetLeft(f.instance)
}

// SetLeft
// CN: 设置左边位置。
// EN: Set Left position.
func (f *TFrame) SetLeft(value int32) {
    Frame_SetLeft(f.instance, value)
}

// Top
// CN: 获取顶边位置。
// EN: Get Top position.
func (f *TFrame) Top() int32 {
    return Frame_GetTop(f.instance)
}

// SetTop
// CN: 设置顶边位置。
// EN: Set Top position.
func (f *TFrame) SetTop(value int32) {
    Frame_SetTop(f.instance, value)
}

// Width
// CN: 获取宽度。
// EN: Get width.
func (f *TFrame) Width() int32 {
    return Frame_GetWidth(f.instance)
}

// SetWidth
// CN: 设置宽度。
// EN: Set width.
func (f *TFrame) SetWidth(value int32) {
    Frame_SetWidth(f.instance, value)
}

// Height
// CN: 获取高度。
// EN: Get height.
func (f *TFrame) Height() int32 {
    return Frame_GetHeight(f.instance)
}

// SetHeight
// CN: 设置高度。
// EN: Set height.
func (f *TFrame) SetHeight(value int32) {
    Frame_SetHeight(f.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (f *TFrame) Cursor() TCursor {
    return Frame_GetCursor(f.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (f *TFrame) SetCursor(value TCursor) {
    Frame_SetCursor(f.instance, value)
}

// Hint
// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (f *TFrame) Hint() string {
    return Frame_GetHint(f.instance)
}

// SetHint
// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (f *TFrame) SetHint(value string) {
    Frame_SetHint(f.instance, value)
}

// Margins
// CN: 获取边矩，仅VCL有效。
// EN: Get Edge moment, only VCL is valid.
func (f *TFrame) Margins() *TMargins {
    return MarginsFromInst(Frame_GetMargins(f.instance))
}

// SetMargins
// CN: 设置边矩，仅VCL有效。
// EN: Set Edge moment, only VCL is valid.
func (f *TFrame) SetMargins(value *TMargins) {
    Frame_SetMargins(f.instance, CheckPtr(value))
}

// CustomHint
// CN: 获取自定义提示。
// EN: Get custom hint.
func (f *TFrame) CustomHint() *TCustomHint {
    return CustomHintFromInst(Frame_GetCustomHint(f.instance))
}

// SetCustomHint
// CN: 设置自定义提示。
// EN: Set custom hint.
func (f *TFrame) SetCustomHint(value IComponent) {
    Frame_SetCustomHint(f.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (f *TFrame) ComponentCount() int32 {
    return Frame_GetComponentCount(f.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (f *TFrame) ComponentIndex() int32 {
    return Frame_GetComponentIndex(f.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (f *TFrame) SetComponentIndex(value int32) {
    Frame_SetComponentIndex(f.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (f *TFrame) Owner() *TComponent {
    return ComponentFromInst(Frame_GetOwner(f.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (f *TFrame) Name() string {
    return Frame_GetName(f.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (f *TFrame) SetName(value string) {
    Frame_SetName(f.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (f *TFrame) Tag() int {
    return Frame_GetTag(f.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (f *TFrame) SetTag(value int) {
    Frame_SetTag(f.instance, value)
}

// DockClients
// CN: 获取指定索引停靠客户端。
// EN: .
func (f *TFrame) DockClients(Index int32) *TControl {
    return ControlFromInst(Frame_GetDockClients(f.instance, Index))
}

// Controls
// CN: 获取指定索引子控件。
// EN: .
func (f *TFrame) Controls(Index int32) *TControl {
    return ControlFromInst(Frame_GetControls(f.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (f *TFrame) Components(AIndex int32) *TComponent {
    return ComponentFromInst(Frame_GetComponents(f.instance, AIndex))
}

