import 'package:chegirma/data/model/categories_response.dart';
import 'package:chegirma/data/repository/main_repository.dart';
import 'package:get/get.dart';

class HomeController extends GetxController {
  bool loading = false;
  List<Categories> category = [];

  @override
  void onInit(){
    super.onInit();
    getCategories();
  }

  getCategories() async {
    loading = true;
    final result = await MainRepository().getCategories();
    loading = false;

    if (result.data is CategoriesResponse) {
      category = result.data?.categories ?? [];
      update();
    } else {
      Get.snackbar("Error", result.toString());
    }
  }
}
