#include <QGLWidget>

class QGLWidgetCustom : public QGLWidget {
	Q_OBJECT // must include this if you use Qt signals/slots

protected slots:
	void initializeGL() {
		emit initializeGLSignal();
	}
	
	void resizeGL(int w, int h) {
		emit resizeGLSignal(w, h);
	}
	
	void paintGL() {
		emit paintGLSignal();
	}

signals:
	void initializeGLSignal();
	void resizeGLSignal(int w, int h);
	void paintGLSignal();
};
