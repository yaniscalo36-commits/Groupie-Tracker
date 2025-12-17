fetch("/api/artists")
  .then(response => {
    if (!response.ok) {
      throw new Error("Erreur HTTP " + response.status);
    }
    return response.json();
  })
  .then(artists => {
    const container = document.getElementById("artists-list");
    const errorMsg = document.getElementById("error-message");
    errorMsg.textContent = ""; // reset message erreur

    container.innerHTML = "";

    artists.forEach(artist => {
      const card = document.createElement("div");
      card.className = "artist-card";

      card.innerHTML = `
        <img src="${artist.image}" alt="${artist.name}">
        <h3>${artist.name}</h3>
        <p>${artist.members.length} membre(s)</p>
      `;

      container.appendChild(card);
    });
  })
  .catch(error => {
    const errorMsg = document.getElementById("error-message");
    errorMsg.textContent = "Erreur lors du chargement des artistes. Veuillez réessayer plus tard.";
    console.error("Erreur API :", error);
  });
