import 'package:flutter/material.dart';

import 'custom_circular_progress_indicator.dart';


class ModalProgressHUD extends StatelessWidget {
  final bool inAsyncCall;
  final double opacity;
  final Color color;
  final Offset? offset;
  final bool dismissible;
  final Widget child;

  const ModalProgressHUD({
    Key? key,
    required this.inAsyncCall,
    this.opacity = 0,
    this.color = Colors.transparent,
    this.offset,
    this.dismissible = false,
    required this.child,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    if (!inAsyncCall) return child;

    Widget layOutProgressIndicator;
    if (offset == null) {
      layOutProgressIndicator = const Center(
        child: CustomCircularProgressIndicator(
          isPage: true,
        ),
      );
    } else {
      layOutProgressIndicator = Positioned(
        child: const CustomCircularProgressIndicator(isPage: true),
        left: offset!.dx,
        top: offset!.dy,
      );
    }

    return Stack(
      children: [
        child,
        Opacity(
          child: ModalBarrier(
            dismissible: dismissible,
            color: color,
          ),
          opacity: opacity,
        ),
        layOutProgressIndicator,
      ],
    );
  }
}
