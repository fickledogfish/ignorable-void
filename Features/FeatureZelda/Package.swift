// swift-tools-version: 6.0

import PackageDescription

let package = Package(
    name: "FeatureZelda",
    platforms: [
        .iOS(.v16),
        .macOS(.v13),
    ],
    products: [
        .library(
            name: "FeatureZelda",
            targets: ["FeatureZelda"]
        ),
    ],
    dependencies: [
        // External
        .package(url: "https://github.com/SimplyDanny/SwiftLintPlugins", exact: "0.57.0"),
        // Feature
        .package(path: "../FeatureSplash"),
    ],
    targets: [
        .target(
            name: "FeatureZelda",
            dependencies: [
                .product(name: "FeatureSplash", package: "FeatureSplash"),
            ],
            path: "Sources",
            plugins: [
                .plugin(name: "SwiftLintBuildToolPlugin", package: "SwiftLintPlugins"),
            ]
        ),
        .testTarget(
            name: "FeatureZeldaTests",
            dependencies: [
                .target(name: "FeatureZelda"),
            ],
            path: "Tests"
        ),
    ],
    swiftLanguageModes: [.v6]
)
