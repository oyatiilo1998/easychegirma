import 'package:chegirma/ui/home_screen.dart';
import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:keyboard_dismisser/keyboard_dismisser.dart';

import 'binding/mian_binding.dart';
import 'core/app_theme.dart';

void main() async {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return KeyboardDismisser(
      gestures: const [GestureType.onTap],
      child: GetMaterialApp(
        builder: (context, child) {
          return MediaQuery(
            data: MediaQuery.of(context).copyWith(
              textScaleFactor: 1.0,
            ),
            child: child!,
          );
        },
        // navigatorKey: ApiClient.alice.getNavigatorKey(),
        theme: appThemeData,
        debugShowCheckedModeBanner: false,
        initialBinding: MainBinding(),
        defaultTransition: Transition.rightToLeft,
        transitionDuration: const Duration(milliseconds: 100),
        home: const HomeScreen(),
      ),
    );
  }
}
