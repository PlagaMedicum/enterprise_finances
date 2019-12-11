import React from "react";
import Modal from "react-bootstrap/Modal";
import Button from "react-bootstrap/Button";

function AddModal(props) {
  const style = {
    color: "black"
  };

  return (
      <Modal show={props.show} onHide={props.hide}>
        <Modal.Header>
          <Modal.Title style={style}>Add Grade</Modal.Title>
        </Modal.Header>
        <Modal.Body style={style}>Body</Modal.Body>
        <Modal.Footer>
          <Button variant="danger" onClick={props.hide}>Close</Button>
        </Modal.Footer>
      </Modal>
  );
}

export default AddModal;