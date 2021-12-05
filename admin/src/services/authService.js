import reqGenerator from "../utils/reqGenerator";

const authService = {
  checkLoginAndPassword: (params) => reqGenerator.post('/user-exists', null, { params })
}

export default authService