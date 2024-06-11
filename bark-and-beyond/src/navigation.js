import React from 'react';
import { Link } from 'react-router-dom'; 

function Navigation() {
  return (
    <nav>
      <ul>
        <li><Link to="/">Home</Link></li>
        <li><Link to="/about">About Us</Link></li> {/* ink to the about us */}
      </ul>
    </nav>
  );
}

export default Navigation;
