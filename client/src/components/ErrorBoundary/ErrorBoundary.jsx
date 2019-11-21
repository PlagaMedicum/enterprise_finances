import React from "react";
import 'bootstrap/dist/css/bootstrap.min.css';
import Alert from 'react-bootstrap/Alert';

class ErrorBoundary extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      error: null,
      errorInfo: null
    };
  }

  componentDidCatch(error, errorInfo) {
    this.setState({error, errorInfo});
  }

  render() {
    if (this.state.errorInfo) {
      return (
          <Alert variant='danger'>
            <Alert.Heading>Error</Alert.Heading>
            {this.state.error && this.state.error.toString()}
            <br/>
            {this.state.errorInfo.componentStack}
          </Alert>
      );
    }

    return this.props.children;
  }
}

export default ErrorBoundary;
