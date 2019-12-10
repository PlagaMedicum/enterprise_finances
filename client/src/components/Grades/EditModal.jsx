import React from "react";
import Modal from "react-bootstrap/Modal";
import Button from "react-bootstrap/Button";

class EditModal extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      show: false
    }
  }

  handleClose() {
    this.setState({show: false})
  }

  render() {
    return (
        <div>
          <Modal show={this.props.show} onHide={this.props.hide}>
            <Modal.Header>
              <Modal.Title style={{color: "black"}}>Header</Modal.Title>
            </Modal.Header>
            <Modal.Body style={{color: "black"}}>Body</Modal.Body>
            <Modal.Footer>
              <Button variant="danger" onClick={this.props.hide}>Close</Button>
            </Modal.Footer>
          </Modal>
        </div>
    );
  }
}

export default EditModal;
