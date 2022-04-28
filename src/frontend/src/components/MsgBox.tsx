import React from 'react'
import { Button, Modal } from 'react-bootstrap'

type Props = {
    show: boolean,
    title: string,
    text: string
    onHide: () => void
}

function MsgBox( { show, title, text, onHide } : Props ) {

  return (
      <Modal
        show={show}
        onHide={onHide}>
            
        <Modal.Header>
            <Modal.Title>{title}</Modal.Title>
        </Modal.Header>
        <Modal.Body>
            { 
                text.split('\n').map((item, index) => {
                    return <p key={index}>{item}</p>
                })
            }
        </Modal.Body>
        <Modal.Footer>
            <Button variant="primary" onClick={onHide}>
                Close
            </Button>
        </Modal.Footer>
    </Modal>
  )
}

export default MsgBox