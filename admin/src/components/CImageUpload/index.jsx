import "./style.scss"

import { CircularProgress } from "@mui/material"
import { useCallback, useRef, useState } from "react"
import { useDropzone } from "react-dropzone"
import UploadIcon from "@mui/icons-material/Upload"
import reqGenerator from "../../utils/reqGenerator"
import CancelIcon from '@mui/icons-material/Cancel';

const CImageUpload = ({value, onChange, error}) => {
  const [loader, setLoader] = useState(false)
  const inputRef = useRef(null)


  const onDrop = useCallback((files) => {
    const uploadedFile = files[0]
    setLoader(true)
    const data = new FormData()
    data.append("image", uploadedFile)

    reqGenerator
      .post("/image-upload", data, {
        headers: {
          "Content-Type": "mulpipart/form-data",
        },
      })
      .then((res) => {
        const url = "https://cdn.ekadastr.udevs.io/ekadastr/" + res.file_path
        onChange(url)
      })
      .finally(() => setLoader(false))
  }, [])

  const { getRootProps, getInputProps } = useDropzone({ onDrop })

  const clearFile = (e) => {
    e.stopPropagation()
    onChange(null)
  }

  if (value) {
    return <div className="image-block">
      <div className="shadow"></div>
      <img
        style={{ width: "100%", height: "100%", objectFit: "cover" }}
        src={value}
        alt=""
      />
      <button className="deleteBtn" onClick={clearFile} ><CancelIcon /></button>
    </div>
  }

  return (
    <div className="FImageUpload">
      <div
        {...getRootProps()}
        className={`dropzone ${error ? "error" : ""}`}
        ref={inputRef}
        style={{ height: 164 }}
      >
        <input {...getInputProps()} />
        {!loader ? (
          <>
            <UploadIcon style={{ fontSize: "50px" }} />
            <p className="title">Faylni yuklang</p>
          </>
        ) : (
          <CircularProgress />
        )}
      </div>
    </div>
  )
}

export default CImageUpload
