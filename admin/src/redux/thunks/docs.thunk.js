import computeActJSON from "../../utils/computeActJSON"
import reqGenerator from "../../utils/reqGenerator"
import { timestamper } from "../../utils/timestamper"
import { docsActions } from "../slices/docs.slice"
import { showAlert } from "./alert.thunk"
import history from "../../history"
import { loaderAction } from "../slices/loader.slice"
import computeEmpowermentJSON from "../../utils/computeEmpowermentJSON"
import computeFacturaJSON from "../../utils/computeFacturaJSON"
import DATA from '../../utils/data.json'
import putSignature from "../../utils/putSignature"

const { loaderOFF, loaderON } = loaderAction

export const fetchMeasureList = () => (dispatch) => {
  reqGenerator.get("/measure/list").then((res) => {
    if (!res) return null
    dispatch(
      docsActions.setMeasureList(
        res.map((el) => ({ value: el.measureId, label: el.name }))
      )
    )
  })
}

export const createActDoc = (values, productList) => async (dispatch, getState) => {
  const guid = await reqGenerator.get("/info/guid")
  const key = getState().auth.key

    let data = null

    try {
      data = JSON.stringify(computeActJSON(values, productList, guid.data))
    } catch (error) {
      return dispatch(showAlert("Malumotlar noto'g'ri kiritilgan"))
    }

    putSignature(data).then(pkcs7 => {
      dispatch(loaderON("fullPageLoader"))
      console.log("PCKS7 ===>", pkcs7)
      reqGenerator
        .post("/docs/act/create", {
            sign: pkcs7,
          },
          { params: { tin: key.TIN } }
        )
        .then((res) => {
          history.push("/main/docs/sender")
          dispatch(
            showAlert("Hujjat muvoffaqiyatli yaratildi", "success")
          )
        })
        .finally(() => dispatch(loaderOFF("fullPageLoader")))
    })
  }

  export const createEmpowermentDoc = (values, productList) => async (dispatch, getState) => {
    const guid = await reqGenerator.get("/info/guid")
    const key = getState().auth.key

    let data = null

    try {
      data = JSON.stringify(computeEmpowermentJSON(values, productList, guid.data))
    } catch (error) {
      return dispatch(showAlert("Malumotlar noto'g'ri kiritilgan"))
    }
    
    console.log("COMPUTED DATA ===>", data)

    putSignature(data).then(pkcs7 => {
      dispatch(loaderON("fullPageLoader"))
      console.log("PCKS7 ===>", pkcs7)
      reqGenerator
        .post("/docs/empowerment/create", {
            sign: pkcs7,
          },
          { params: { tin: key.TIN } }
        )
        .then((res) => {
          history.push("/main/docs/sender")
          dispatch(
            showAlert("Hujjat muvoffaqiyatli yaratildi", "success")
          )
        })
        .finally(() => dispatch(loaderOFF("fullPageLoader")))
    })

    // let data = null
    // try {
    //   data = JSON.stringify(computeEmpowermentJSON(values, productList, guid))
    // } catch (error) {
    //   console.log("error => ", error)
    //   return dispatch(showAlert("Malumotlar noto'g'ri kiritilgan"))
    // }
    // dispatch(loaderON("fullPageLoader"))

    // dispatch(showAlert("Kalit uchun parolni kiriting", "info"))
    // await window.EIMZOClient.loadKey(
    //   key,
    //   (keyId) => {
    //     window.EIMZOClient.createPkcs7(
    //       keyId,
    //       data,
    //       timestamper,
    //       async (pkcs7) => {
    //         console.log("PCK7 ==> ", pkcs7)
    //         reqGenerator
    //           .post(
    //             "/docs/empowerment/create",
    //             {
    //               sign: pkcs7,
    //             },
    //             { params: { tin: key.TIN } }
    //           )
    //           .then((res) => {
    //             history.push("/main/docs/seller")
    //             dispatch(
    //               showAlert("Hujjat muvoffaqiyatli yaratildi", "success")
    //             )
    //           })
    //           .finally(() => dispatch(loaderOFF("fullPageLoader")))
    //       },
    //       () => {
    //         dispatch(showAlert("Kalit uchun parolni noto'g'ri kiritdingiz"))
    //         dispatch(loaderOFF("fullPageLoader"))
    //       }
    //     )
    //   },
    //   () => {
    //     dispatch(showAlert("Error"))
    //     dispatch(loaderOFF("fullPageLoader"))
    //   }
    // )
  }



  export const createFacturaDoc = (values, productList) => async (dispatch, getState) => {
    const { guid } = await reqGenerator.get("/info/guid")
    const key = getState().auth.key

    let data = null
    try {
      data = JSON.stringify(computeFacturaJSON(values, productList, guid))
    } catch (error) {
      console.log("error => ", error)
      return dispatch(showAlert("Malumotlar noto'g'ri kiritilgan"))
    }
    dispatch(loaderON("fullPageLoader"))

    dispatch(showAlert("Kalit uchun parolni kiriting", "info"))
    await window.EIMZOClient.loadKey(
      key,
      (keyId) => {
        window.EIMZOClient.createPkcs7(
          keyId,
          data,
          timestamper,
          async (pkcs7) => {
            console.log("PCK7 ==> ", pkcs7)
            reqGenerator
              .post(
                "/docs/factura/create",
                {
                  sign: pkcs7,
                },
                { params: { tin: key.TIN } }
              )
              .then((res) => {
                history.push("/main/docs/sender")
                dispatch(
                  showAlert("Hujjat muvoffaqiyatli yaratildi", "success")
                )
              })
              .finally(() => dispatch(loaderOFF("fullPageLoader")))
          },
          () => {
            dispatch(showAlert("Kalit uchun parolni noto'g'ri kiritdingiz"))
            dispatch(loaderOFF("fullPageLoader"))
          }
        )
      },
      () => {
        dispatch(showAlert("Error"))
        dispatch(loaderOFF("fullPageLoader"))
      }
    )
  }
