## work in progress

## TODO

- [ ] UI
  - [ ] direct spi vs device tree + framebuffer
  - [ ] layout
  - [ ] only use touchscreen, or hw button to wake??
  - [ ] backlight control?
  - [ ] with dt+framebuffer, how to get touch events?
- [ ] EAP requirements
  - [ ] RADIUS server
        - customize - detect activity and extend acct lifetime?
- [ ] touchscreen interface
- [ ] generate qr code with ssid, device-specific password, etc
  - [ ] ban device if MAC changes? (might mess with anti-tracking/privacy features on device...)
- [ ] alternate for human input into devices that don't read wifi qr codes
  - [ ] simpler passwords?
  - [ ] longer expiry?
  - [ ] record mode in database for future lookup
  - [ ] require person's name??
  - [ ] detect multiple devices sharing a password - allow, deny, flag?
- [ ] 
- [ ] 
- [ ] 
- [ ] 
- [ ] 

## SO discussion
https://stackoverflow.com/a/60605927/382458
> I found some information on how to format the WiFi config string in the following pull request at the github page of the zxing library project: https://github.com/zxing/zxing/pull/865
> 
> The first post contains a template of the string format, including an error (the prefix AI: is wrong, it must read A:, see here). The correct format according to the source is thus:
> 
> WIFI:T:WPA2-EAP;S:[network SSID];H:[hidden?];E:[EAP method];PH2:[Phase 2 method];A:[anonymous identity];I:[username];P:[password];;
When I tried this (using the command line tool qrencode) my Barcode Scanner app crashed. After some trial and error I figured that the option for hiding the SSID can be left out:
> 
> WIFI:T:WPA2-EAP;S:[network SSID];E:[EAP method];PH2:[Phase 2 method];A:[anonymous identity];I:[username];P:[password];;
With this I'm getting a working entry in the list of known wireless networks in Android 8.
> 
> As of now there is no support for declaring a certificate and the respective domain. If this is needed, one can specify it later by adjusting the settings from inside Android's WiFi menu.

> Nice to see it is possible. I had the thought of using EAP on my home network and setting up a mini kiosk (esp32 + lcd?) with a button. Pressing the button would cause an account to be created, with a QR code displayed on-screen for connection of the smartphone. This means a guest can connect to wifi without much trouble. I assume it wouldn't be too difficult to remove/deactivate accounts in RADIUS when they are inactive for a period of time, limiting the number of devices in the wild containing valid credentials for the network.

> I had the same idea, just with a raspberry pi I'm using anyway. You could specify the expiration date of a new generated RADIUS user to be 2 days in the future. That way the automatically generated accounts would expire without you having to worry about anything.

## protocols
https://www.securew2.com/blog/wpa2-enterprise-authentication-protocols-comparison
* EAP-TLS sounds best but requires certs which apparently don't work with QR codes
* EAP-TTLS/PAP uses passwords, should be compatible with QR
* MSCHAP is _bad_
* other options??
* what does the AP support?

## packages

* radius
  * layeh.com/radius
