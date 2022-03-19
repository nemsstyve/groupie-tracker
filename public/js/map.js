//Étape 1: On initialise la communication avec la plateforme
//remplacez la variable window.apikey par votre propre apikey
var platform = new H.service.Platform({
    apikey: "CcbbjF82rHrjrBx7E6xPP7gZvGKVC15ALYkzp7zgg84"
});
var defaultLayers = platform.createDefaultLayers();

//Étape 2: On initialiser une carte - cette carte est centrée sur l'Europe
var map = new H.Map(document.getElementById('map'),
    defaultLayers.vector.normal.map, {
        center: { lat: 50, lng: 5 },
        zoom: 3,
        pixelRatio: window.devicePixelRatio || 1
    });
//On ajoute un écouteur de redimensionnement pour s'assurer que la carte occupe tout le conteneur
window.addEventListener('resize', () => map.getViewPort().resize());

//Étape 3 : rendre la carte interactive
//MapEvents active le système d'événements
//Le comportement implémente les interactions par défaut pour le panoramique/zoom (également sur les environnements tactiles mobiles)
var behavior = new H.mapevents.Behavior(new H.mapevents.MapEvents(map));

//On Crée les composants d'interface utilisateur par défaut
var ui = H.ui.UI.createDefault(map, defaultLayers);

//On Utilise maintenant la carte au besoin...
window.onload = async function() {
    let markers = document.getElementsByClassName("locations");
    for (let i = 0; i < markers.length; i++) {
        const data = await fetch(`https://geocoder.ls.hereapi.com/6.2/geocode.json?searchtext=${markers[i].textContent}&gen=9&apiKey=CcbbjF82rHrjrBx7E6xPP7gZvGKVC15ALYkzp7zgg84`)
        const response = await data.json()
        const lat = response.Response.View[0].Result[0].Location.DisplayPosition.Latitude
        const lng = response.Response.View[0].Result[0].Location.DisplayPosition.Longitude
        map.addObject(new H.map.Marker({ lat, lng }));
    }
}