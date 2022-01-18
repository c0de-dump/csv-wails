import React, { useCallback } from 'react'
import { useDropzone } from 'react-dropzone'

function FileDropZone() {
  const onDrop = useCallback((file) => {})

  const { getRootProps, getInputProps, isDragActive } = useDropzone({ onDrop })

  return (
    <section>
      <div {...getRootProps({ className: 'dropzone' })}>
        <input {...getInputProps()} />
        <p> Drag 'n' Drop </p>
      </div>
    </section>
  )
}
