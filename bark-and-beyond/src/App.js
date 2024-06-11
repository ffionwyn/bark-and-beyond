import React from 'react';
import { BrowserRouter as Router, Route } from 'react-router-dom';
import Navigation from './navigation.js';
import AboutUs from './pages/about-us';

function App() {
  return (
    <Router>
      <div>
        <Navigation />
        <Route path="/about">
          <AboutUs />
        </Route>
        {/*add more routes */}
      </div>
    </Router>
  );
}

export default App;
