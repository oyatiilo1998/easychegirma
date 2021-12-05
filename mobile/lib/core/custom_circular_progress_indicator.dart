import 'dart:math' as math;

import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';

import 'app_theme.dart';

class CustomCircularProgressIndicator extends StatefulWidget {
  final bool isPage;

  const CustomCircularProgressIndicator({
    Key? key,
    this.isPage = false,
  }) : super(key: key);

  @override
  _CustomCircularProgressIndicatorState createState() => _CustomCircularProgressIndicatorState();
}

class _CustomCircularProgressIndicatorState extends State<CustomCircularProgressIndicator> with SingleTickerProviderStateMixin {
  late AnimationController? _controller;

  @override
  void initState() {
    super.initState();
    _controller = AnimationController(
      vsync: this,
      duration: const Duration(
        milliseconds: 800,
      ),
    )..repeat();
  }

  @override
  Widget build(BuildContext context) {
    return Center(
      child: widget.isPage
          ? AnimatedBuilder(
          animation: _controller!.view,
          builder: (context, snapshot) {
            return Transform.rotate(
              angle: _controller!.value * 2 * math.pi,
              child: Container(
                height: 48,
                width: 48,
                decoration: const BoxDecoration(
                  shape: BoxShape.circle,
                  color: second,
                ),
                child: Center(
                  child: Column(
                    crossAxisAlignment: CrossAxisAlignment.center,
                    mainAxisAlignment: MainAxisAlignment.center,
                    children: [
                      Row(
                        crossAxisAlignment: CrossAxisAlignment.center,
                        mainAxisAlignment: MainAxisAlignment.center,
                        children: [
                          circleWidget(),
                          const SizedBox(
                            width: 4,
                          ),
                          circleWidget(),
                        ],
                      ),
                      const SizedBox(
                        height: 4,
                      ),
                      Row(
                        crossAxisAlignment: CrossAxisAlignment.center,
                        mainAxisAlignment: MainAxisAlignment.center,
                        children: [
                          circleWidget(),
                          const SizedBox(
                            width: 4,
                          ),
                          circleWidget(),
                        ],
                      ),
                    ],
                  ),
                ),
              ),
            );
          })
          : const CircularProgressIndicator(),
    );
  }

  Widget circleWidget() {
    return Container(
      width: 12,
      height: 12,
      decoration: const BoxDecoration(
        color: first,
        shape: BoxShape.circle,
      ),
    );
  }

  @override
  void dispose() {
    _controller?.dispose();
    super.dispose();
  }
}
