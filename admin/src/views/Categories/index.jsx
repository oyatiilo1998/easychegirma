import { Button } from "@mui/material";
import { useHistory } from "react-router";
import Header from "../../components/Header";
import Table from "./Table";

const CategoryList = () => {
  const history = useHistory()

  return (
    <div>
      <Header title="Список категорий">
        <Button color="secondary" variant="contained" onClick={() => history.push("/main/categories/create")} >
          Создат новый
        </Button>
      </Header>
      <div
        style={{ padding: "20px", display: "flex", flexDirection: "column" }}
      >
        <Table />
      </div>
    </div>
  );
};

export default CategoryList;
