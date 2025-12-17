fetch("/api/artists")
  .then(res => {
    if (!res.ok) {
      throw new Error("Erreur HTTP " + res.status);
    }
    return res.json();
  })
  .then(artists => {
    const container = document.getElementById("trending-list");
    container.innerHTML = "";

    artists.slice(0, 8).forEach(artist => {
      container.innerHTML += `
        <div class="artist-card">
          <img src="${artist.image}" alt="${artist.name}">
          <h3>${artist.name}</h3>
        </div>
      `;
    });
  })
  .catch(error => {
    console.error("Erreur API :", error);
  });
