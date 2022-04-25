import React from 'react'
import { Button, Form } from 'react-bootstrap'

function DNATest() {
    const [name, setName] = React.useState('')
    const [sequence, setSequence] = React.useState('')
    const [disease, setDisease] = React.useState('')
    const [status, setStatus] = React.useState('')
    const [selected, setSelected] = React.useState('bm')

    const handleChangeFile = (event: React.ChangeEvent<HTMLInputElement>) => {
        if (event.target.files) {
            const file = event.target.files[0];
            const reader = new FileReader();
            reader.readAsText(file);
            reader.onload = () => {
                setSequence(reader.result as string);
            }
        }
    }

    const handleOptionChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setSelected(event.target.value)
    }

    const handleSubmit = async (e: React.MouseEvent<HTMLButtonElement>) => {
        e.preventDefault();
        try {
            const response = await fetch('/api/match/'+selected, {
                method: 'POST',
                mode: 'same-origin',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    "name": name,
                    "sequence": sequence,
                    "penyakit": disease
                })
            })
            if(response.ok){
                setStatus('Success')
            } else {
                setStatus(await response.text())
            }
        } catch (error) {
            setStatus('Internal Server error')
        }
    }

  return (
    <div className = 'container mt-5'>
        <div className='card'>
            <h1>DNA Test</h1>
            <p>Status: {status}</p>
            <Form className='form'>
                <p className='mt-2'>Username</p>
                <input type="text" placeholder="Name" onChange={e => setName(e.target.value)}/>
                <p>Sequence</p>
                <input type="file" onChange={handleChangeFile}/>
                <p>Disease</p>
                <input type="text" placeholder='Disease' onChange={e => setDisease(e.target.value)}/>
                <label>
                  <input type="radio" value="bm" checked={selected === 'bm'} onChange={handleOptionChange} />
                  Boyer Moore (BM)
                </label>
                <label>
                  <input type="radio" value="kmp" checked={selected === 'kmp'} onChange={handleOptionChange} />
                  Knuth Morris Pratt (KMP)
                </label>
                <Button variant="primary" onClick={handleSubmit} className='button mt-2'>Submit</Button>
            </Form>
        </div>
    </div>
  )
}

export default DNATest