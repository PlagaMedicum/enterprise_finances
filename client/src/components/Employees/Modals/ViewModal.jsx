import React from "react";
import Modal from "react-bootstrap/Modal";
import Button from "react-bootstrap/Button";

function ViewModal(props) {
  const style = {
    color: "black"
  };

  return (
      <Modal style={style} show={props.show} onHide={props.hide}>
        <Modal.Header>
          <Modal.Title>View Employee</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <b>#{props.id}</b><br/>
          <b>Name:</b> {}<br/>
          <b>Position:</b> {}
        </Modal.Body>
        <Modal.Footer>
          <Button variant="danger" onClick={props.hide}>Close</Button>
        </Modal.Footer>
      </Modal>
  );
}

export default ViewModal;
