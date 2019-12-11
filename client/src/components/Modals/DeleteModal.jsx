import React from "react";
import Modal from "react-bootstrap/Modal";
import Button from "react-bootstrap/Button";

function DeleteModal(props) {
  const style = {
    color: "black"
  };

  return (
      <Modal style={style} show={props.show} onHide={props.hide}>
        <Modal.Header>
          <Modal.Title>{props.title}</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          Do you really want to delete this piece of very secret information?<br/>
          Note: Don't ever make those steps without agreement of your superior. Elsewise you will be threaten as hostile
          subject which must be immediately destroyed with flame and blame!
        </Modal.Body>
        <Modal.Footer>
          <Button variant="success" onClick={props.confirm}>Agree</Button>
          <Button variant="danger" onClick={props.hide}>Close</Button>
        </Modal.Footer>
      </Modal>
  );
}

export default DeleteModal;
