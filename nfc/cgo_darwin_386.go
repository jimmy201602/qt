package nfc

/*
#cgo CPPFLAGS: -pipe -fpascal-strings -fmessage-length=0 -Wno-trigraphs -Wreturn-type -Wparentheses -Wswitch -Wno-unused-parameter -Wunused-variable -Wunused-value -Wno-shorten-64-to-32 -Wno-sign-conversion -fexceptions -fasm-blocks -Wno-missing-field-initializers -Wno-missing-prototypes -Wno-implicit-atomic-properties -Wformat -Wno-missing-braces -Wno-unused-function -Wno-unused-label -Wuninitialized -Wno-unknown-pragmas -Wno-shadow -Wno-four-char-constants -Wno-sign-compare -Wpointer-sign -Wno-newline-eof -Wdeprecated-declarations -Winvalid-offsetof -Wno-conversion -O2 -isysroot /Applications/Xcode.app/Contents/Developer/Platforms/iPhoneSimulator.platform/Developer/SDKs/iPhoneSimulator9.3.sdk -mios-simulator-version-min=6.1 -arch i386 -fobjc-nonfragile-abi -fobjc-legacy-dispatch -Wno-deprecated-implementations -Wprotocol -Wno-selector -Wno-strict-selector-match -Wno-undeclared-selector -Wall -W -fPIC
#cgo CPPFLAGS: -DDARWIN_NO_CARBON -DQT_NO_PRINTER -DQT_NO_PRINTDIALOG -DQT_NO_DEBUG -DQT_CORE_LIB -DQT_NFC_LIB
#cgo CPPFLAGS: -I/usr/local/Qt5.6.0/5.6/ios/mkspecs/macx-ios-clang/ios -I/usr/local/Qt5.6.0/5.6/ios/include -I/usr/local/Qt5.6.0/5.6/ios/mkspecs/macx-ios-clang
#cgo CPPFLAGS: -I/usr/local/Qt5.6.0/5.6/ios/include/QtCore -I/usr/local/Qt5.6.0/5.6/ios/include/QtNfc

#cgo LDFLAGS: -headerpad_max_install_names -stdlib=libc++ -Wl,-syslibroot,/Applications/Xcode.app/Contents/Developer/Platforms/iPhoneSimulator.platform/Developer/SDKs/iPhoneSimulator9.3.sdk -mios-simulator-version-min=6.1 -arch i386
#cgo LDFLAGS: -L/usr/local/Qt5.6.0/5.6/ios/plugins/platforms -lqios_iphonesimulator -framework Foundation -framework UIKit -framework QuartzCore -framework AssetsLibrary -L/usr/local/Qt5.6.0/5.6/ios/lib -framework MobileCoreServices -framework CoreFoundation -framework CoreText -framework CoreGraphics -framework OpenGLES -lqtfreetype_iphonesimulator -framework Security -framework SystemConfiguration -framework CoreBluetooth -L/usr/local/Qt5.6.0/5.6/ios/plugins/imageformats -lqdds_iphonesimulator -lqicns_iphonesimulator -lqico_iphonesimulator -lqtga_iphonesimulator -lqtiff_iphonesimulator -lqwbmp_iphonesimulator -lqwebp_iphonesimulator -lqtharfbuzzng_iphonesimulator -lz -lqtpcre_iphonesimulator -lm -lQt5Core_iphonesimulator -lQt5Nfc_iphonesimulator -lQt5Gui_iphonesimulator -lQt5PlatformSupport_iphonesimulator
*/
import "C"
