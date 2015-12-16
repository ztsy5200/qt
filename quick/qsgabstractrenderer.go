package quick

//#include "quick.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QSGAbstractRenderer struct {
	core.QObject
}

type QSGAbstractRenderer_ITF interface {
	core.QObject_ITF
	QSGAbstractRenderer_PTR() *QSGAbstractRenderer
}

func PointerFromQSGAbstractRenderer(ptr QSGAbstractRenderer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGAbstractRenderer_PTR().Pointer()
	}
	return nil
}

func NewQSGAbstractRendererFromPointer(ptr unsafe.Pointer) *QSGAbstractRenderer {
	var n = new(QSGAbstractRenderer)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QSGAbstractRenderer_") {
		n.SetObjectName("QSGAbstractRenderer_" + qt.Identifier())
	}
	return n
}

func (ptr *QSGAbstractRenderer) QSGAbstractRenderer_PTR() *QSGAbstractRenderer {
	return ptr
}

//QSGAbstractRenderer::ClearModeBit
type QSGAbstractRenderer__ClearModeBit int64

const (
	QSGAbstractRenderer__ClearColorBuffer   = QSGAbstractRenderer__ClearModeBit(0x0001)
	QSGAbstractRenderer__ClearDepthBuffer   = QSGAbstractRenderer__ClearModeBit(0x0002)
	QSGAbstractRenderer__ClearStencilBuffer = QSGAbstractRenderer__ClearModeBit(0x0004)
)

func (ptr *QSGAbstractRenderer) ClearColor() *gui.QColor {
	defer qt.Recovering("QSGAbstractRenderer::clearColor")

	if ptr.Pointer() != nil {
		return gui.NewQColorFromPointer(C.QSGAbstractRenderer_ClearColor(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGAbstractRenderer) ClearMode() QSGAbstractRenderer__ClearModeBit {
	defer qt.Recovering("QSGAbstractRenderer::clearMode")

	if ptr.Pointer() != nil {
		return QSGAbstractRenderer__ClearModeBit(C.QSGAbstractRenderer_ClearMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGAbstractRenderer) DeviceRect() *core.QRect {
	defer qt.Recovering("QSGAbstractRenderer::deviceRect")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QSGAbstractRenderer_DeviceRect(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGAbstractRenderer) ConnectSceneGraphChanged(f func()) {
	defer qt.Recovering("connect QSGAbstractRenderer::sceneGraphChanged")

	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_ConnectSceneGraphChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sceneGraphChanged", f)
	}
}

func (ptr *QSGAbstractRenderer) DisconnectSceneGraphChanged() {
	defer qt.Recovering("disconnect QSGAbstractRenderer::sceneGraphChanged")

	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_DisconnectSceneGraphChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sceneGraphChanged")
	}
}

//export callbackQSGAbstractRendererSceneGraphChanged
func callbackQSGAbstractRendererSceneGraphChanged(ptrName *C.char) {
	defer qt.Recovering("callback QSGAbstractRenderer::sceneGraphChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "sceneGraphChanged")
	if signal != nil {
		signal.(func())()
	}

}

func (ptr *QSGAbstractRenderer) SetClearColor(color gui.QColor_ITF) {
	defer qt.Recovering("QSGAbstractRenderer::setClearColor")

	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_SetClearColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *QSGAbstractRenderer) SetClearMode(mode QSGAbstractRenderer__ClearModeBit) {
	defer qt.Recovering("QSGAbstractRenderer::setClearMode")

	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_SetClearMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QSGAbstractRenderer) SetDeviceRect(rect core.QRect_ITF) {
	defer qt.Recovering("QSGAbstractRenderer::setDeviceRect")

	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_SetDeviceRect(ptr.Pointer(), core.PointerFromQRect(rect))
	}
}

func (ptr *QSGAbstractRenderer) SetDeviceRect2(size core.QSize_ITF) {
	defer qt.Recovering("QSGAbstractRenderer::setDeviceRect")

	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_SetDeviceRect2(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QSGAbstractRenderer) SetProjectionMatrix(matrix gui.QMatrix4x4_ITF) {
	defer qt.Recovering("QSGAbstractRenderer::setProjectionMatrix")

	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_SetProjectionMatrix(ptr.Pointer(), gui.PointerFromQMatrix4x4(matrix))
	}
}

func (ptr *QSGAbstractRenderer) SetProjectionMatrixToRect(rect core.QRectF_ITF) {
	defer qt.Recovering("QSGAbstractRenderer::setProjectionMatrixToRect")

	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_SetProjectionMatrixToRect(ptr.Pointer(), core.PointerFromQRectF(rect))
	}
}

func (ptr *QSGAbstractRenderer) SetViewportRect(rect core.QRect_ITF) {
	defer qt.Recovering("QSGAbstractRenderer::setViewportRect")

	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_SetViewportRect(ptr.Pointer(), core.PointerFromQRect(rect))
	}
}

func (ptr *QSGAbstractRenderer) SetViewportRect2(size core.QSize_ITF) {
	defer qt.Recovering("QSGAbstractRenderer::setViewportRect")

	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_SetViewportRect2(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QSGAbstractRenderer) ViewportRect() *core.QRect {
	defer qt.Recovering("QSGAbstractRenderer::viewportRect")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QSGAbstractRenderer_ViewportRect(ptr.Pointer()))
	}
	return nil
}