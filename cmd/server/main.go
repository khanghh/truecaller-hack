package main

import (
	"fmt"
	"runtime"

	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

var (
	AppName    string
	Version    string
	CommitHash string
	BuiltTime  string
	OsArch     string
)

const hostName = "truecaller.xoac.xyz"

var installlationInfo interface{}

func init() {
	AppName = "TrueCaller Server"
	OsArch = runtime.GOOS + "/" + runtime.GOARCH
}

func printVersion() {
	fmt.Println(AppName)
	fmt.Printf(" Version:\t%s\n", Version)
	fmt.Printf(" Commit:\t%s\n", CommitHash)
	fmt.Printf(" Built Time:\t%s\n", BuiltTime)
	fmt.Printf(" OS/Arch:\t%s\n", OsArch)
}

func serveProxyAPI(router *gin.RouterGroup, apiUrl *url.URL) {
	router.Any("/*proxyPath", func(c *gin.Context) {
		proxy := httputil.NewSingleHostReverseProxy(apiUrl)
		proxy.Director = func(req *http.Request) {
			req.Header = c.Request.Header
			req.Host = apiUrl.Host
			req.URL.Scheme = apiUrl.Scheme
			req.URL.Host = apiUrl.Host
			req.URL.Path = c.Param("proxyPath")
		}
		proxy.ServeHTTP(c.Writer, c.Request)
	})
}

func handleSendOnboardingOtp(c *gin.Context) {
	jsonStr := `{
		"status": 1,
		"message": "Sent",
		"domain": "noneu",
		"parsedPhoneNumber": 84328289006,
		"parsedCountryCode": "VN",
		"requestId": "916fcd3d-5c76-405b-b315-d3138c3bb5d1",
		"method": "sms",
		"tokenTtl": 300
	}`

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, jsonStr)
}

func handleVerifyOnboardingOtp(c *gin.Context) {
	jsonStr := `{
		"status": 2,
		"message": "Verified",
		"installationId": "a1i0q--X-Qylb-JVy_pheZhAXIqjyF6EZ3x73gOBqQ4CUWkQEUszlh1JY6DAHMnd",
		"ttl": 259200,
		"userId": 9572072038335776,
		"suspended": false,
		"phones": [{
			"phoneNumber": 84328289006,
			"countryCode": "VN",
			"priority": 1
		}]
	}`

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, jsonStr)
}

func handleGetConfiguration(c *gin.Context) {
	jsonStr := `{
		"features": {
			"featureSearchMissTtl": "86400000",
			"featureDetailsAdsRemovalForPriorityAndVb": "1",
			"featureFetchingTopSpammersHoursInterval": "24",
			"featureContextCall": "1",
			"featureDualNumberEditProfile": "1",
			"featureWhatsAppCalls": "1",
			"featurePushCallerId": "0",
			"featureNewDetailsViewForSpammers": "1",
			"threeButtonPremiumLayoutEnabled": "1",
			"featureShowContactTimezone": "1",
			"featureInsightsHideTrxAction": "1",
			"featureResourceBrokenUpdate": "1",
			"featureCampaignKeywordsOnPrefs": "1",
			"featureGoldPremiumGift": "1",
			"featureWhatsappCallerId": "1",
			"featureImNewJoinersPeriodDays": "7",
			"featureCallSilencePushCallerId": "1",
			"featureOptimizedAdsNativeView": "1",
			"featureBusinessProfiles": "1",
			"featureWhoViewedMe": "1",
			"featureFlash": "1",
			"featureEnableTagsInACS": "1",
			"featureBackup": "1",
			"featureShowComments": "1",
			"featureCallLogTcx": "1",
			"featureDisabledExtendedPrivacy": "1",
			"featureFileMimeTypes": "application/vnd.ms-powerpoint|application/vnd.openxmlformats-officedocument.presentationml.presentation|application/vnd.ms-excel|application/vnd.openxmlformats-officedocument.spreadsheetml.sheet|application/msword|application/vnd.openxmlformats-officedocument.wordprocessingml.document|application/pdf|text/plain",
			"featureCacheAdAfterCall": "1",
			"featureBizCallReasonForBusinesses": "1",
			"featureLargeDetailsViewAd": "1",
			"featureDuo": "0",
			"featureScheduleMessage": "1",
			"featureCallRecordingsScopedStorageMigration": "1",
			"featurePromoIncomingMsgCount": "1",
			"featureNameFeedbackCooldown": "1",
			"featureSearchWarnings": "1",
			"featureVisiblePushCallerId": "1",
			"featureImPromoAfterCallPeriodDays": "0",
			"featureNetworkConnection": "1",
			"featureNameSuggestions": "1",
			"featureInsightsUserFeedback": "1",
			"featureAcsAdsRemovalForPriorityAndVb": "1",
			"featureFullscreenACS": "1",
			"featureNewDetailsViewAll": "1",
			"featureRemoveInvalidCallLogEntries": "1",
			"featureGetImUserTtl": "89280000",
			"featureAdRetentionTime": "5",
			"featureImTracingEnabled": "0",
			"featureInCallUIDefaultOptIn": "1",
			"featureImGroupMaxParticipantCount": "200",
			"featureHMSAttestation": "1",
			"featureGetImUserTtlConversation": "3600000",
			"featureDisplaySpamStats": "1",
			"featureLogEcommerceEvents": "0",
			"featureExitPhoneStateChangeEarlyOnRinging": "1",
			"featureYearInReview_v2021": "1",
			"featureEnableGoldCallerIdForContacts": "1",
			"featureReferralNavDrawer": "1",
			"featureInsightsAnalytics": "1",
			"featureReferralDeeplink": "1",
			"featureCrossDcSearch": "1",
			"featureGroupInviteLinks": "1",
			"featureTagUpload": "1000",
			"featureGetImUserMissTtl": "3600000",
			"featureLinkPreviews": "1",
			"featureRichTextFormatting": "1",
			"featureContactAboutAsPremium": "1",
			"featureCallRecordingAccessibility": "1",
			"featureAddFeedbackCommentBox": "1",
			"featureShowACSPbSetting": "1",
			"featureAdsInterstitial": "0",
			"featureIncognitoOnProfileViewCount": "6",
			"featureAds": "0",
			"featureDisableBusinessImCategorization": "0",
			"featureLocationPreview": "1",
			"featureAppsInstalledHeartbeat": "1",
			"featureIdleCallStateMonitor": "1",
			"featureReactionsEmojis": "\uD83D\uDC4D,\uD83E\uDD23,\uD83D\uDE2E,\uD83D\uDE0D,\uD83D\uDE20,\uD83D\uDE22,\uD83D\uDC4E",
			"featureSearchThrottlingHandler": "1",
			"featurePresenceInterval": "15000",
			"featureBlockOptionsClevertap": "1",
			"featureTruecallerNewsMenuItem": "1",
			"featureGroupAutoJoin": "1",
			"featureGlobalSearchRevamp": "1",
			"featureAfterBlockPromo": "1",
			"featurePresenceInitialDelay": "500",
			"featureOutgoingSearch": "1",
			"featureIMReply": "1",
			"featureLoggingEnabled": "0",
			"featureConversationSpamSearchCount": "150",
			"featureWVMWeeklySummaryNotification": "1",
			"featureHouseAdsTimeout": "10",
			"featureSmartNotifications": "1",
			"featureNameFeedbackCooldownPeriodSeconds": "86400",
			"featureImEmojiPoke": "1",
			"featureAdPartnerSdkMediation": "1",
			"featureDetailsViewPullToRefresh": "1",
			"featureMarkAsImportantROW": "1",
			"featureInsights": "1",
			"featureAdUnitIdCache": "1",
			"featureUploadEventsJitter": "10000",
			"redesignEditProfile": "1",
			"featureAnnounceCallerId": "1",
			"featureShowUserJoinedImNotification": "1",
			"featureShareImageInFlash": "1",
			"featurePersonalSafetyPromo": "1",
			"featureBusinessIm": "1",
			"featureAskDisableBatteryOptimization": "1",
			"featureBizNewDetailsViewVerifiedBusinessProfiles": "1",
			"featureSdkBottomSheetDialog": "1",
			"featureVerifiedBusinessAwareness": "1",
			"featurePresenceWithoutJobScheduler": "1",
			"featureNumberScanner": "1",
			"featureSearchBarAnimation": "1",
			"featureReferralIconInSearch": "1",
			"featureTruecallerX": "1",
			"featureManageDataRegion2": "1",
			"featureVoteComments": "1",
			"featureTopSpammersPremiumMaxSize": "40000",
			"featureInsightsRerun": "1",
			"featureEUp": "1",
			"featureReferralAfterCallSaveContact": "1",
			"featureContactJobAsPremium": "1",
			"featurePremiumUserTab": "1",
			"featureHideACSSetting": "1",
			"featurePlacesSDK": "1",
			"featureAcsRateAppPromo": "1",
			"featureBizVideoCallerId": "1",
			"featureSdk1tap": "1",
			"featureShowACSAllOutgoing": "1",
			"featureWhoViewedMeShowNotificationAfterXLookups": "1",
			"featureForcedUpdateDialog": "1",
			"featureSearchInConversation": "1",
			"featureInsightsUserFeedbackButton": "1",
			"featureAdRouterOnGAM": "1",
			"featurePresenceStopTime": "600000",
			"featureVoiceClip": "1",
			"featureMaxEventsBatchSize": "100",
			"featureGhostCall": "1",
			"featureContactFieldsPremiumForProfile": "1",
			"featureDetailsViewNewCacheBehaviour": "1",
			"featureFiveBottomTabsWithBlockingPremium": "1",
			"featureOtpParserRegex": "(?i)(?:(?:otp|verification|verify|pin|password|code).*?(?:\\\\s+is\\\\s+|\\\\s+is:\\\\s*|(?:a-zA-z:-\\\\s)*))([0-9]{3,8})/(?i)(?:^(?:enter|use)*.*?)([0-9]{4,8})(?:(?:\\\\s+(?:is|as|to)\\\\s.+?)(?:otp|verification|verify|pin|password|code))",
			"featureDeviceAttestation": "1",
			"featureOtpNotification": "1",
			"featureVoipAfterCallPromoDays": "5",
			"featureWhitelistedFBLogNotifications": "1",
			"featureTracingEnabled": "0",
			"featureAdsUnifiedBlock": "1",
			"featureNewACS": "1",
			"featureContactEmailAsPremium": "1",
			"featureNameFeedbackInterval": "7200",
			"featureAttachmentPickerV2": "1",
			"featureFBLogNotifications": "0",
			"featureVideoCallerId": "1",
			"featureCreateBusinessProfiles": "1",
			"featureShowLargeBannerAdsOnAftercall": "1",
			"featureBlockReasons": "1",
			"featureReferralShareApps": "App Chooser",
			"featureHideDialpad": "1",
			"featureCallRecordingNewDesign": "1",
			"featureDisplaySpamCategories": "1",
			"featureVoipRedesign": "1",
			"featureCrossDcVoip": "1",
			"featureAllowSortComments": "1",
			"featureInCallUI": "1",
			"featureNudgeToSendAsSMS": "1",
			"featureNewCallsFragment": "1",
			"featureLogCallEventsV3": "1",
			"featureBusinessSuggestion": "1",
			"featureEnableNewNativeAdImageTemplate": "1",
			"featureCommentsDefaultSortByScore": "1",
			"featureWhoViewedMeACSEnabled": "1",
			"featureInCallUIWhatsNew": "1",
			"featureCleverTap": "1",
			"featureFetchBusinessCardOnPremiumStatusChange": "1",
			"featureSyncPhonebookOnCallStateChange": "1",
			"featureAcsDetectPhonebookContactsNewBehaviour": "1",
			"featureVoIP": "1",
			"featurePushCallerIdV2": "1",
			"featureStatsSearch": "1",
			"featureV2TaggerSearchUi": "1",
			"featureGAMInternalEvent": "1",
			"featureVoIPGroup": "1",
			"featureContactsListV2": "1",
			"featureOtpConversationSmartAction": "1",
			"featureUrlMinSpamScore": "0",
			"featureTCY": "1",
			"featureCallingRedesignDetails": "1",
			"featureEnableMediumBannerACS": "1",
			"featurePersonalSafetyMenuItem": "1",
			"defaultSMSAppPromoDuration": "28",
			"featureWhoViewedMeIncognito": "Pro",
			"featureAlternativeDau": "1",
			"featureSearchWarningFeatureStore": "1",
			"featureShowACSAllIncoming": "1",
			"featueNewAdsKeywords": "1",
			"featureInCallUISwitchToVoip": "1",
			"featureSupernovaOptOutVisible": "1",
			"featurePredictiveECPMModel": "1",
			"featureCrossDomainPresence": "1",
			"featureAdsUnifiedDetails": "1",
			"featureSearchHitTtl": "1296000000",
			"featureExtendedT9Search": "1",
			"featurePresenceRecheckTime": "13",
			"featureWhoViewedMeShowNotificationAfterXDays": "1",
			"featureCallIntentPackage": "1",
			"featurePromoSpamOffCount": "1",
			"defaultSMSAppPromoDate": "Aug 23, 2021",
			"featureMinEventsBatchSize": "100",
			"featureUploadComments": "1",
			"featureContactAddressAsPremium": "1",
			"featureFBLogBackgroundWork": "0",
			"featureWhoViewedMeNewViewIntervalInDays": "1",
			"featureImMaxMediaSize": "104857600",
			"featurePromotionalMessageCategory": "1",
			"featureInsightsReconciliation": "1",
			"featureIMLocationPreviewParameters": "480|320|17",
			"featureConvertBusinessProfileToPrivate": "1",
			"feature121MutingAndSounds": "1",
			"featureReactions": "1",
			"featureExitPhoneStateChangeEarlyIfStateMismatch": "1",
			"featureShowRingingDuration": "1",
			"searchVersion": "1",
			"featureAvailability": "1",
			"featureIm": "1",
			"featureCacheOnInCallNotification": "1",
			"featureContactSocialAsPremium": "1",
			"featureNameFeedback": "1",
			"featureContactFieldsPremiumForUgc": "1"
		}
	}`
	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, jsonStr)
}

func handleGetSubscribtions(c *gin.Context) {
	jsonStr := `{
		"proStatus": {
			"credits": 0,
			"start": "1970-01-01T00:00:00Z",
			"expires": "2970-01-01T00:00:00Z",
			"gracePeriodExpires": "2970-01-01T00:00:00Z",
			"renewable": false,
			"level": 100,
			"isFreeTrial": false,
			"isExpired": false,
			"isGracePeriodExpired": false,
			"contactRequestQuota": {
				"month": 30,
				"Year": 365
			},
			"inAppPurchaseAllowed": true
		}
	}`
	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, jsonStr)
}

func handleEdgeLocations(c *gin.Context) {
	jsonTmp := `{
		"data": {
		  "noneu": {
			"phonebook5": {
			  "edges": [
				"phonebook5-noneu.truecaller.com"
			  ]
			},
			"referrals": {
			  "edges": [
				"referrals-noneu.truecaller.com"
			  ]
			},
			"backup": {
			  "edges": [
				"backup-noneu.truecaller.com"
			  ]
			},
			"pushid": {
			  "edges": [
				"pushid-noneu.truecaller.com"
			  ]
			},
			"callmeback": {
			  "edges": [
				"callmeback-noneu.truecaller.com"
			  ]
			},
			"callkit": {
			  "edges": [
				"callkit-noneu.truecaller.com"
			  ]
			},
			"edge-locations5": {
			  "edges": [
				"edge-locations5-noneu.truecaller.com"
			  ]
			},
			"topspammers": {
			  "edges": [
				"topspammers-noneu.truecaller.com"
			  ]
			},
			"ads5": {
			  "edges": [
				"ads5-noneu.truecaller.com"
			  ]
			},
			"availability5": {
			  "edges": [
				"availability5.truecaller.com"
			  ]
			},
			"feedback": {
			  "edges": [
				"feedback-noneu.truecaller.com"
			  ]
			},
			"profile4": {
			  "edges": [
				"profile4-noneu.truecaller.com"
			  ]
			},
			"profile-view": {
			  "edges": [
				"profile-view-noneu.truecaller.com"
			  ]
			},
			"search5": {
			  "edges": [
				"search5-noneu.truecaller.com"
			  ]
			},
			"link-reports": {
			  "edges": [
				"link-reports-noneu.truecaller.com"
			  ]
			},
			"opt-out": {
			  "edges": [
				"opt-out-asia-south1.truecaller.com"
			  ]
			},
			"messenger": {
			  "edges": [
				"messenger-noneu.truecaller.com"
			  ]
			},
			"tagging5": {
			  "edges": [
				"tagging5-noneu.truecaller.com"
			  ]
			},
			"voip": {
			  "edges": [
				"voip-noneu.truecaller.com"
			  ]
			},
			"device-safety": {
			  "edges": [
				"device-safety-noneu.truecaller.com"
			  ]
			},
			"ads-router": {
			  "edges": [
				"ads-router-noneu.truecaller.com"
			  ]
			},
			"lastactivity": {
			  "edges": [
				"lastactivity-noneu.truecaller.com"
			  ]
			},
			"presence": {
			  "edges": [
				"presence.truecaller.com"
			  ]
			},
			"contact-request": {
			  "edges": [
				"contact-request-noneu.truecaller.com"
			  ]
			},
			"premium": {
			  "edges": [
				"%s"
			  ]
			},
			"notifications5": {
			  "edges": [
				"notifications5-noneu.truecaller.com"
			  ]
			},
			"ads-segment": {
			  "edges": [
				"ads-segment-profile-noneu.truecaller.com"
			  ]
			},
			"oauth-account": {
			  "edges": [
				"oauth-account-noneu.truecaller.com"
			  ]
			},
			"userapps": {
			  "edges": [
				"userapps-noneu.truecaller.com"
			  ]
			},
			"batchlogging4": {
			  "edges": [
				"batchlogging4-noneu.truecaller.com"
			  ]
			},
			"leadgen": {
			  "edges": [
				"leadgen-noneu.truecaller.com"
			  ]
			},
			"flash": {
			  "edges": [
				"flash-noneu.truecaller.com"
			  ]
			},
			"account": {
			  "edges": [
				"%s"
			  ]
			},
			"filter-store4": {
			  "edges": [
				"filter-store4-noneu.truecaller.com"
			  ]
			},
			"contact-upload4": {
			  "edges": [
				"contact-upload4-noneu.truecaller.com"
			  ]
			},
			"api4": {
			  "edges": [
				"api4-noneu.truecaller.com"
			  ]
			},
			"user-archive": {
			  "edges": [
				"user-archive-noneu.truecaller.com"
			  ]
			},
			"images": {
			  "edges": [
				"images-noneu.truecaller.com"
			  ]
			},
			"presence-grpc": {
			  "edges": [
				"presence-grpc-noneu.truecaller.com"
			  ]
			},
			"pstn": {
			  "edges": [
				"pstn-noneu.truecaller.com"
			  ]
			}
		  },
		  "eu": {
			"phonebook5": {
			  "edges": [
				"phonebook5-eu.truecaller.com"
			  ]
			},
			"referrals": {
			  "edges": [
				"referrals-eu.truecaller.com"
			  ]
			},
			"backup": {
			  "edges": [
				"backup-eu.truecaller.com"
			  ]
			},
			"pushid": {
			  "edges": [
				"pushid-eu.truecaller.com"
			  ]
			},
			"callmeback": {
			  "edges": [
				"callmeback-eu.truecaller.com"
			  ]
			},
			"callkit": {
			  "edges": [
				"callkit-eu.truecaller.com"
			  ]
			},
			"lastactivity": {
			  "edges": [
				"lastactivity-eu.truecaller.com"
			  ]
			},
			"edge-locations5": {
			  "edges": [
				"edge-locations5-eu.truecaller.com"
			  ]
			},
			"topspammers": {
			  "edges": [
				"topspammers-eu.truecaller.com"
			  ]
			},
			"ads5": {
			  "edges": [
				"ads5-eu.truecaller.com"
			  ]
			},
			"availability5": {
			  "edges": [
				"availability5.truecaller.com"
			  ]
			},
			"feedback": {
			  "edges": [
				"feedback-eu.truecaller.com"
			  ]
			},
			"profile4": {
			  "edges": [
				"profile4-eu.truecaller.com"
			  ]
			},
			"profile-view": {
			  "edges": [
				"profile-view-eu.truecaller.com"
			  ]
			},
			"userapps": {
			  "edges": [
				"userapps-eu.truecaller.com"
			  ]
			},
			"search5": {
			  "edges": [
				"search5-eu.truecaller.com"
			  ]
			},
			"link-reports": {
			  "edges": [
				"link-reports-eu.truecaller.com"
			  ]
			},
			"opt-out": {
			  "edges": [
				"opt-out-eu.truecaller.com"
			  ]
			},
			"push-callerid": {
			  "edges": [
				"push-callerid-eu.truecaller.com"
			  ]
			},
			"messenger": {
			  "edges": [
				"messenger-eu.truecaller.com"
			  ]
			},
			"tagging5": {
			  "edges": [
				"tagging5-eu.truecaller.com"
			  ]
			},
			"voip": {
			  "edges": [
				"voip-eu.truecaller.com"
			  ]
			},
			"device-safety": {
			  "edges": [
				"device-safety-eu.truecaller.com"
			  ]
			},
			"presence": {
			  "edges": [
				"presence-eu.truecaller.com"
			  ]
			},
			"contact-request": {
			  "edges": [
				"contact-request-eu.truecaller.com"
			  ]
			},
			"premium": {
			  "edges": [
				"premium-se1.truecaller.com"
			  ]
			},
			"notifications5": {
			  "edges": [
				"notifications5-eu.truecaller.com"
			  ]
			},
			"batchlogging4": {
			  "edges": [
				"batchlogging4-eu.truecaller.com"
			  ]
			},
			"leadgen": {
			  "edges": [
				"leadgen-eu.truecaller.com"
			  ]
			},
			"flash": {
			  "edges": [
				"flash-eu.truecaller.com"
			  ]
			},
			"account": {
			  "edges": [
				"account-eu.truecaller.com"
			  ]
			},
			"filter-store4": {
			  "edges": [
				"filter-store4-eu.truecaller.com"
			  ]
			},
			"contact-upload4": {
			  "edges": [
				"contact-upload4-eu.truecaller.com"
			  ]
			},
			"api4": {
			  "edges": [
				"api4-eu.truecaller.com"
			  ]
			},
			"user-archive": {
			  "edges": [
				"user-archive-eu.truecaller.com"
			  ]
			},
			"images": {
			  "edges": [
				"images-eu.truecaller.com"
			  ]
			},
			"presence-grpc": {
			  "edges": [
				"presence-grpc-eu.truecaller.com"
			  ]
			},
			"pstn": {
			  "edges": [
				"pstn-eu.truecaller.com"
			  ]
			}
		  }
		},
		"ttl": 43200
	  }`
	jsonStr := fmt.Sprintf(jsonTmp, hostName)
	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, jsonStr)
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/v2", handleEdgeLocations)
	router.POST("/v2/sendOnboardingOtp", handleSendOnboardingOtp)
	router.GET("/v2/subscriptions", handleGetSubscribtions)
	router.POST("/v1/verifyOnboardingOtp", handleVerifyOnboardingOtp)
	router.PUT("/v1/installation")
	router.GET("/v1/configuration", handleGetConfiguration)

	router.Run(":3000")
}
