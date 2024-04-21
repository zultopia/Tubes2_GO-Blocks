import React from 'react';

function Navbar() {
  return (
    <nav className="nav">
        <div>
            <ul>
                <li><a href="/home">Home</a></li>
                <li><a href="/howtouse">How To Use</a></li>
                <li><a href="/about">About Us</a></li>
            </ul>
        </div>
    </nav>
  );
}

export default Navbar;