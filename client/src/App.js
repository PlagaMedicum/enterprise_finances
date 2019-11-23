import React, {Component} from 'react';
import {BrowserRouter as Router, Route, Switch} from "react-router-dom";
import ErrorBoundary from "./components/ErrorBoundary/ErrorBoundary";
import NavBar from "./components/NavBar/NavBar";
import Home from "./components/Home/Home";
import Employees from "./components/Employees/Employees";
import Grades from "./components/Grades/Grades";

class App extends Component {
  render() {
    const style = {
      padding: 50
    };

    return (
        <Router>
          <div className="App">
            <ErrorBoundary>
              <NavBar/>
            </ErrorBoundary>
            <ErrorBoundary>
              <section style={style}>
                <Switch>
                  <Route path="/" exact component={Home}/>
                  <Route path="/employees" component={Employees}/>
                  <Route path="/grades" component={Grades}/>
                </Switch>
              </section>
            </ErrorBoundary>
          </div>
        </Router>
    );
  }
}

export default App;
