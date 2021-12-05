import 'package:cached_network_image/cached_network_image.dart';
import 'package:chegirma/controller/home_controller.dart';
import 'package:chegirma/core/modal_progress_hud.dart';
import 'package:chegirma/ui/home_carusel.dart';
import 'package:flutter/material.dart';
import 'package:get/get.dart';

import 'category_screen.dart';

class HomeScreen extends GetView<HomeController> {
  const HomeScreen({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return GetBuilder<HomeController>(builder: (logic) {
      return Scaffold(
        appBar: AppBar(
          backgroundColor: Colors.white,
          toolbarHeight: 0,
          centerTitle: true,
        ),
        body: ModalProgressHUD(
          inAsyncCall: logic.loading,
          child: ListView(
            children: [
              const HomeCarousel(),
              for (int i = 0; i < logic.category.length; i++)
                Padding(
                  padding: const EdgeInsets.only(top: 16, left: 16, right: 16),
                  child: InkWell(
                    onTap: () {
                      Get.to(CategoryScreen(logic.category[i]));
                    },
                    borderRadius: BorderRadius.circular(16),
                    child: Ink(
                      height: 144,
                      decoration: const BoxDecoration(color: Color(0xFFF4f4f4), borderRadius: BorderRadius.all(Radius.circular(16))),
                      child: Row(
                        mainAxisAlignment: MainAxisAlignment.center,
                        children: [
                          Padding(
                            padding: const EdgeInsets.only(left: 12),
                            child: Text(
                              logic.category[i].name ?? '',
                              style: const TextStyle(
                                fontSize: 22,
                                fontWeight: FontWeight.w700,
                              ),
                            ),
                          ),
                          const Spacer(),
                          Padding(
                            padding: const EdgeInsets.all(12.0),
                            child: ClipRRect(
                              borderRadius: BorderRadius.circular(12),
                              child: CachedNetworkImage(
                                imageUrl: logic.category[i].ruName ?? '',
                                height: 144,
                                width: 144,
                                alignment: Alignment.bottomLeft,
                                fit: BoxFit.cover,
                                progressIndicatorBuilder: (_, __, ___) {
                                  return SizedBox();
                                },
                              ),
                            ),
                          ),
                        ],
                      ),
                    ),
                  ),
                ),
              const SizedBox(
                height: 16,
              )
            ],
          ),
        ),
      );
    });
  }
}
