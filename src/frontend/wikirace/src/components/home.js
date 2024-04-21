import React, {useState} from 'react';
import './styles.css';
import wikirace from './assets/wikirace.png'
import stima from './assets/stima.png'
import goblocks from './assets/goblocks.png'
import bfs_text from './assets/bfs.png'
import ids_text from './assets/ids.png'
import './styles.css';
import mcqueen from './assets/mcqueen.png';
import cruz from './assets/cruz.png';
import start from './assets/start.png';
import end from './assets/end.png';
import search from './assets/search.png';

function Home() {
  const [startArticle, setStartArticle] = useState('');
  const [targetArticle, setTargetArticle] = useState('');
  const [result, setResult] = useState(null);
  const [isLoading, setIsLoading] = useState(false); // Add loading state

  const handleSubmit = async (e) => {
    e.preventDefault();
    setIsLoading(true);
    try {
        // Construct full URLs
        const fullStartArticleURL = `https://en.wikipedia.org/wiki/${startArticle}`;
        const fullTargetArticleURL = `https://en.wikipedia.org/wiki/${targetArticle}`;

        // Make API request with full URLs
        const response = await fetch(`http://localhost:8080/shortestpath?start=${encodeURIComponent(fullStartArticleURL)}&target=${encodeURIComponent(fullTargetArticleURL)}`);
        const data = await response.json();
        setResult(data);
        setIsLoading(false);
        console.log('Data fetched successfully:', data);
    } catch (error) {
        console.error('Error fetching data:', error);
        setIsLoading(false);
    }
  };
  
  return (
    <div className="home-background">
      {/* Gambar dan tombol */}
      <div>
        <img src={wikirace} alt="Wiki Race" className="wiki-race-image" />
        <img src={stima} alt="STIMA" className="stima-image" />
        <img src={goblocks} alt="GO Blocks" className="goblocks-image" />
        <img src={bfs_text} alt="BFS Title" className="bfs-text" />
        <img src={ids_text} alt="IDS Title" className="ids-text" />
        <img src={start} alt="START" className="start-image" />
        <img src={end} alt="END" className="end-image" />
        <img src={search} alt="SEARCH" className="search-image" />
        <button className="mcqueen-button">
          <img src={mcqueen} alt="Button 1" />
        </button>
        <button className="cruz-button">
          <img src={cruz} alt="Button 2" />
        </button>
      </div>

      {/* Kontainer Logika */}
      <div className="logic-container">
        <div className="start-container">
          <input
            type="text"
            value={startArticle}
            onChange={(e) => setStartArticle(e.target.value)}
            placeholder="Start Article"
          />
        </div>

        <div class="end-container">
          <input
            type="text"
            value={targetArticle}
            onChange={(e) => setTargetArticle(e.target.value)}
            placeholder="End Article"
          />
        </div>

        <div className="search-container">
          <button onClick={handleSubmit}>Search</button>
        </div>

        {/* Hasil Pencarian */}
        <div className="result-container">
          {isLoading ? (
            <p>Loading...</p>
          ) : result ? (
            <div>
              <h2>Result</h2>
              <ul>
                {result.path.map((url, index) => (
                  <li key={index}>
                    <a href={url} target="_blank" rel="noopener noreferrer">
                      {url}
                    </a>
                  </li>
                ))}
              </ul>
              <p>Articles Visited: {result.articlesVisited}</p>
              <p>Articles Checked: {result.articlesChecked}</p>
              <p>Execution Time: {result.executionTime} ms</p>
            </div>
          ) : null}
        </div>
      </div>
    </div>
  );
}

export default Home;