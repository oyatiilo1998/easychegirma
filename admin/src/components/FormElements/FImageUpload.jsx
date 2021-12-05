import CImageUpload from "../CImageUpload"
import _ from 'lodash'

const FImageUpload = ({ formik, name, ...props }) => {


  return (
    <CImageUpload
      name={name}
      value={_.get(formik?.values, name)}
      onChange={val => formik.setFieldValue(name, val)}
      error={_.get(formik.touched, name) && Boolean(_.get(formik.errors, name))}
      {...props}
    />
  )
}

export default FImageUpload
