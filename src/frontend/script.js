document.getElementById('wikiRaceForm').addEventListener('submit', async (e) => {
    e.preventDefault();
    
    const startArticle = document.getElementById('startArticle').value;
    const endArticle = document.getElementById('endArticle').value;
    
    const response = await fetch(`/search?start=${startArticle}&end=${endArticle}`);
    const data = await response.json();
    
    if (data.path) {
        const resultDiv = document.getElementById('result');
        resultDiv.innerHTML = `<p>Number of articles checked: ${data.nodesChecked}</p>`;
        resultDiv.innerHTML += `<p>Number of articles traversed: ${data.path.length}</p>`;
        resultDiv.innerHTML += `<p>Traversal route:</p>`;
        resultDiv.innerHTML += `<ul>`;
        data.path.forEach(page => {
            resultDiv.innerHTML += `<li>${page.Title}</li>`;
        });
        resultDiv.innerHTML += `</ul>`;
        resultDiv.innerHTML += `<p>Search time: ${data.searchTime} ms</p>`;
    } else {
        alert('No path found.');
    }
});