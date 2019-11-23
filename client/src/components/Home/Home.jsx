import React from "react";
import Alert from "react-bootstrap/Alert";
import Image from "react-bootstrap/Image";
import header from "./img/header.png"
import ResponsiveEmbed from "react-bootstrap/ResponsiveEmbed";
import Jumbotron from "react-bootstrap/Jumbotron";

class Home extends React.Component {
  render() {
    const style = {
      textAlign: "center",
      backgroundColor: "#111111"
    };

    return (
        <div>
          <header style={style}>
            <Image src={header} fluid/>
          </header>
          <Alert variant="warning">
            <Alert.Heading>Welcome to the Black Mesa Financial Systems!</Alert.Heading>
            <p>
              Do you know who ate all the donuts?
            </p>
            <hr/>
            <p>
              Why do we all must to wear those ridiculous ties?!
            </p>
          </Alert>
          <Jumbotron>
            <ResponsiveEmbed aspectRatio="21by9">
              <iframe className="embed-responsive-item" src="https://www.youtube.com/embed/BJTAzUArRTo?rel=0"
                      allowFullScreen/>
            </ResponsiveEmbed>
            <br/>
            <Alert variant="success">
              Ah ello there, Gordn Freeman!
            </Alert>
          </Jumbotron>

        </div>
    );
  }
}

export default Home;
