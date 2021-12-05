import 'package:chegirma/data/model/products_response.dart';
import 'package:chegirma/data/repository/main_repository.dart';
import 'package:get/get.dart';

class CategoryController extends GetxController {
  bool loading = false;
  List<Products> products = [];

  getProducts(String id) async {
    loading = true;
    final result = await MainRepository().getProducts(categoryId: id);
    loading = false;

    if (result.data is ProductsResponse) {
      products = result.data?.products ?? [];
      update();
    } else {
      Get.snackbar("Error", result.toString());
    }
  }
}
