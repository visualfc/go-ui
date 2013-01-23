#include <QGLWidget>

class QGLWidgetCustom : public QGLWidget {
	Q_OBJECT // must include this if you use Qt signals/slots

protected slots:
	void initializeGL() {
		emit initializeGLSignal();
	}
	
	void resizeGL(int w, int h) {
		emit resizeGLSignal(w,h);
	}
	
	void paintGL() {
		emit paintGLSignal();
	}
	
	void initializeOverlayGL() {
		emit initializeOverlayGLSignal();
	}
	
	void resizeOverlayGL(int w, int h) {
		emit resizeOverlayGLSignal(w,h);
	}
	
	void paintOverlayGL() {
		emit paintOverlayGLSignal();
	}

signals:
	void initializeGLSignal();
	void resizeGLSignal(int,int);
	void paintGLSignal();
	void initializeOverlayGLSignal();
	void resizeOverlayGLSignal(int,int);
	void paintOverlayGLSignal();
};
