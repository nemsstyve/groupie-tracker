//Step 1: initialize communication with the platform
// In your own code, replace variable window.apikey with your own apikey
var platform = new H.service.Platform({
    apikey: "CcbbjF82rHrjrBx7E6xPP7gZvGKVC15ALYkzp7zgg84"
});
var defaultLayers = platform.createDefaultLayers();

//Step 2: initialize a map - this map is centered over Europe
var map = new H.Map(document.getElementById('map'), 
    defaultLayers.vector.normal.map,{
    center: {lat:50, lng:5},
    zoom: 3,
    pixelRatio: window.devicePixelRatio || 1
});
// add a resize listener to make sure that the map occupies the whole container
window.addEventListener('resize', () => map.getViewPort().resize());

//Step 3: make the map interactive
// MapEvents enables the event system
// Behavior implements default interactions for pan/zoom (also on mobile touch environments)
var behavior = new H.mapevents.Behavior(new H.mapevents.MapEvents(map));

// Create the default UI components
var ui = H.ui.UI.createDefault(map, defaultLayers);

// Now use the map as required...
window.onload = async function () {
    let markers = document.getElementsByClassName("locations");
    for (let i = 0; i < markers.length; i++) {
        const data = await fetch(`https://geocoder.ls.hereapi.com/6.2/geocode.json?searchtext=${markers[i].textContent}&gen=9&apiKey=CcbbjF82rHrjrBx7E6xPP7gZvGKVC15ALYkzp7zgg84`)
        const response = await data.json()

        const lat = response.Response.View[0].Result[0].Location.DisplayPosition.Latitude
        const lng = response.Response.View[0].Result[0].Location.DisplayPosition.Longitude
        map.addObject(new H.map.Marker({lat, lng}));
    }
}