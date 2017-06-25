import QtQuick 2.8
import QtPositioning 5.5
import QtLocation 5.6

Rectangle {

    property alias routesModel: routesView.model

    function navigate(latlng) {
        var gcoord = _toGeoCoordinate(latlng)
        map.center = gcoord;
        map.zoomLevel = 14
    }

    function _toGeoCoordinate(latlng) {
        return QtPositioning.coordinate(latlng[0], latlng[1])
    }

    Plugin {
        id: osmPlugin
        name: "osm"
        PluginParameter { name: "osm.mapping.host"; value: "http://b.tiles.maps.sputnik.ru/" }
        PluginParameter { name: "osm.mapping.providersrepository.disabled"; value: true }

        //PluginParameter { name: "osm.mapping.host"; value: "http://b.osm.maptiles.xyz/" }
        //PluginParameter { name: "osm.mapping.providersrepository.disabled"; value: true }
        //PluginParameter { name: "osm.mapping.providersrepository.address"; value: "cycle" }

    }


    Map {
        id: map
        anchors.fill: parent
        plugin: osmPlugin
        activeMapType:  mapType(MapType.CustomMap)
        zoomLevel: 13

        //tilt: 50
        //bearing: 90

        center {
            latitude: 54.7514134
            longitude: 20.512562
        }

        function mapType(style) {
            for (var i in map.supportedMapTypes) {
                var t = map.supportedMapTypes[i]
                if (t.style === style) {
                    return t
                }
            }
            return undefined
        }


        MapItemView {
            id: routesView
            delegate: MapPolyline {
                line.width: 4
                line.color: item.color
                path: item.path
            }
        }
    }
}
