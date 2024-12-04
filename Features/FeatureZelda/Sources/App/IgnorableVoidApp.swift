import protocol SwiftUI.App
import protocol SwiftUI.Scene
import struct SwiftUI.WindowGroup

public struct IgnorableVoidApp {
    // MARK: - Initializers

    public init() {
    }
}

// MARK: - App

extension IgnorableVoidApp: App {
    public var body: some Scene {
        WindowGroup {
            ContentView()
        }
    }
}
