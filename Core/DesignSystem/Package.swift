// swift-tools-version: 6.0

import PackageDescription

let package = Package(
    name: "DesignSystem",
    platforms: [
        .macOS(.v12),
        .iOS(.v16),
    ],
    products: [
        .library(
            name: "DesignSystem",
            targets: ["DesignSystem"]
        ),
    ],
    dependencies: [
        // External
        .package(url: "https://github.com/SimplyDanny/SwiftLintPlugins", exact: "0.57.0"),
    ],
    targets: [
        .target(
            name: "DesignSystem",
            path: "Sources",
            plugins: [
                .plugin(name: "SwiftLintBuildToolPlugin", package: "SwiftLintPlugins"),
            ]
        ),
        .testTarget(
            name: "DesignSystemTests",
            dependencies: [
                .target(name: "DesignSystem"),
            ],
            path: "Tests"
        ),
    ],
    swiftLanguageModes: [.v6]
)
