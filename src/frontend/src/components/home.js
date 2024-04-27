import React, {useState} from 'react';
import './styles.css';
import wikirace from './assets/wikirace.png'
import stima from './assets/stima.png'
import goblocks from './assets/goblocks.png'
import bfs_text from './assets/bfs.png'
import ids_text from './assets/ids.png'
import bfs_text2 from './assets/bfs2.png'
import ids_text2 from './assets/ids2.png'
import './styles.css';
import mcqueen from './assets/mcqueen.png';
import cruz from './assets/cruz.png';
import start from './assets/start.png';
import end from './assets/end.png';
import search from './assets/search.png';
import piston from './assets/piston.png';
import switchs from './assets/switch.png';
import WikipediaAutosuggest from './wikipediaautosuggest';
import WikipediaAutosuggest2 from './wikipediaautosuggest2';
import Graf from './graph'
import Switch from './switch';
import one from './assets/one.png';
import many from './assets/many.png';
import kachow from './assets/kachow.gif';

function Home() {
  const [startArticle, setStartArticle] = useState('');
  const [targetArticle, setTargetArticle] = useState('');
  const [result, setResult] = useState(null);
  const [isLoading, setIsLoading] = useState(false); 
  const [bfsSrc, setBfsSrc] = useState(bfs_text); 
  const [idsSrc, setIdsSrc] = useState(ids_text); 
  const [isStartAutocompleteOpen, setIsStartAutocompleteOpen] = useState(false);
  const [isEndAutocompleteOpen, setIsEndAutocompleteOpen] = useState(false);
  const [isMulti, setMulti] = useState(false)

  const handleMcQueenClick = () => {
    setBfsSrc((prevSrc) => (prevSrc === bfs_text ? bfs_text2 : bfs_text)); 
    setIdsSrc(ids_text); 
  };

  const handleCruzClick = () => {
    setIdsSrc((prevSrc) => (prevSrc === ids_text ? ids_text2 : ids_text)); 
    setBfsSrc(bfs_text); 
  };

  const handleSwitch = () => {
    const temp = startArticle;
    setStartArticle(targetArticle);
    setTargetArticle(temp);
  };

  const handleStartSelect = (suggestion) => {
    setStartArticle(suggestion);
    setIsStartAutocompleteOpen(false); 
  };

  const handleEndSelect = (suggestion) => {
    setTargetArticle(suggestion);
    setIsEndAutocompleteOpen(false); 
  };

  /*
  const handleStartContainerClick = () => {
    if (startArticle.length > 0) {
      setIsStartAutocompleteOpen(true); 
    } else {
      setIsStartAutocompleteOpen(false); 
    }
  }; */

  /*
  const handleEndContainerClick = () => {
    if (targetArticle.length > 0) {
      setIsEndAutocompleteOpen(true); 
    } else {
      setIsEndAutocompleteOpen(false); 
    }
  }; */

  const handleStartInputChange = (e) => {
    setStartArticle(e.target.value);
    setIsStartAutocompleteOpen(true); 
  };

  const handleEndInputChange = (e) => {
    setTargetArticle(e.target.value);
    setIsEndAutocompleteOpen(true); 
  };

  const handleSwitchToggle = (isOn) => {
    if (isOn) {
      setMulti(true)
    } else {
      setMulti(false)
    }
  };
  
  const handleSubmit = async (e) => {
    e.preventDefault();
    setIsLoading(true);
    try {
        // Construct full URLs
        const s = startArticle.replaceAll(" ", "_");
        const t = targetArticle.replaceAll(" ", "_");
        const fullStartArticleURL = `https://en.wikipedia.org/wiki/${s}`;
        const fullTargetArticleURL = `https://en.wikipedia.org/wiki/${t}`;
        const Multi = isMulti ? "1" : "2";

        // Make API request with full URLs
        const response = await fetch(`/api/${bfsSrc === bfs_text2 ? "BFS" : "IDS"}`, {
          method: 'POST',
          headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            startTitle: s,
            startURL: fullStartArticleURL,
            targetTitle: t,
            targetURL: fullTargetArticleURL,
            multi: Multi,
          })
        });
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
        <img src={bfsSrc} alt="BFS Title" className="bfs-text" /> 
        <img src={idsSrc} alt="IDS Title" className="ids-text" /> 
        <img src={start} alt="START" className="start-image" />
        <img src={end} alt="END" className="end-image" />
        <img src={one} alt="ONE" className="one-image" />
        <img src={many} alt="MANY" className="many-image" />
        <img src={search} alt="SEARCH" className="search-image" />
        <button className="switch-button">
          <img src={switchs} alt="Switch" onClick={handleSwitch} />
        </button>
        <button className="mcqueen-button" onClick={handleMcQueenClick}>
          <img src={mcqueen} alt="Button 1" />
        </button>
        <button className="cruz-button" onClick={handleCruzClick}>
          <img src={cruz} alt="Button 2" />
        </button>
        <Switch onToggle={handleSwitchToggle} />
        <button className="piston-button" onClick={handleSubmit}>
            <img src={piston} alt="Piston" />
          </button>
      </div>

      {/* Kontainer Logika */}
      <div className="logic-container">
        <div className="start-container">
          <input
            type="text"
            value={startArticle}
            onChange={handleStartInputChange}
            className="start-container"
            placeholder="Start Article"
          />
          {isStartAutocompleteOpen && (
            <WikipediaAutosuggest
              value={startArticle}
              onSelect={handleStartSelect}
              close={() => setIsStartAutocompleteOpen(false)}
          />
        )}
        </div>

        <div className="end-container">
          <input
            type="text"
            value={targetArticle}
            onChange={handleEndInputChange}
            className="end-container"
            placeholder="End Article"
          />
          {isEndAutocompleteOpen && (
            <WikipediaAutosuggest2
              value={targetArticle}
              onSelect={handleEndSelect}
              close={() => setIsEndAutocompleteOpen(false)}
            />
          )}
        </div>

        {/* Hasil Pencarian */}
        <div className="result-container">
          {isLoading ? (
          <div>
            {'Loading...'.split('').map((char, index) => (
              <span key={index} className="loading-text">{char}</span>
            ))}
            <img src={kachow} alt="KACHOW"/>
          </div>
            ) : result ? (
            <div className='graf-wrapper'>
              <Graf 
                paths = {result.paths}
              />
            </div>
          ) : null}
        </div>

        {/* Informasi Pencarian */}
        <div className="info-container">
          {isLoading ? (
            <div>
              {'Loading...'.split('').map((char, index) => (
                <span key={index} className="loading-text">{char}</span>
              ))}
            </div>
          ) : result ? (
            <div className="result-content"> 
              <h2>Result</h2>
              <p>Jumlah Artikel yang diperiksa: {result.articlesVisited}</p>
              <p>Jumlah Artikel yang dilalui: {result.paths.length === 0 ? "No Solution" : result.paths[0].length}</p>
              <p>Execution Time: {result.executionTime} ms</p>
              {result.paths.map((path, index) => (
                <div key={index}>
                <h3>Path {index + 1}</h3>
                  <ul>
                    {path.map((page, pageIndex) => (
                      <li key={pageIndex}>
                        <a href={page.URL} target="_blank" rel="noopener noreferrer">
                          {page.Title}
                        </a>
                      </li>
                    ))}
                </ul>
              </div>
            ))}
            </div>
          ) : null}
          </div>
      </div>
    </div>
  );
}

export default Home;