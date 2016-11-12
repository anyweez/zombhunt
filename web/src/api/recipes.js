export function SearchRecipes(fragment, cb) {
    fetch(`http://localhost:8080/recipes?q=${fragment}`)
        .then(response => response.json())
        .then(response => {
            cb(response);
        });
}