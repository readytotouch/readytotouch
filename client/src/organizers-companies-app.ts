function markCompanyFavorite(companyId: number, favorite: boolean) {
    fetch(`/api/v1/companies/${companyId}/favorite.json`, {
        method: "PATCH",
        body: JSON.stringify({
            favorite: favorite,
        }),
    }).then(console.log).catch(console.error);
}

function markCompanyFavoriteAll(favoriteGenerator: () => boolean) {
    document.querySelectorAll(".js-company").forEach(function ($element: HTMLElement) {
        markCompanyFavorite(parseInt($element.getAttribute("data-company-id")), favoriteGenerator());
    })
}

markCompanyFavoriteAll(function () {
    return Math.random() > 0.5;
});
