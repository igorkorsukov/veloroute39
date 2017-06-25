import QtQuick 2.8
import QtQuick.Window 2.2
import QtQuick.Controls 2.2
import QtQuick.Layouts 1.2

Window {
    id: window
    visible: true
    width: 640
    height: 480
    title: "veloroute39 2"


    Rectangle {
        anchors.fill: parent
        color: "#92B363"
    }

    MapView {
        anchors.fill: parent
    }
}
