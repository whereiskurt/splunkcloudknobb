package service

// SplunkDashboardDataUIView is what is returned from the "Dashboard" view
type SplunkDashboardDataUIView struct {
	Generator struct {
		Build   string `json:"build"`
		Version string `json:"version"`
	} `json:"generator"`

	Links struct {
		Create string `json:"create"`
		Reload string `json:"_reload"`
		ACL    string `json:"_acl"`
	} `json:"links"`

	Origin  string `json:"origin"`
	Updated string `json:"updated"`

	Entry []struct {
		Name    string `json:"name"`
		ID      string `json:"id"`
		Author  string `json:"author"`
		Updated string `json:"updated"`

		Links struct {
			Alternate string `json:"alternate"`
			List      string `json:"list"`
			Reload    string `json:"_reload"`
			Edit      string `json:"edit"`
			Remove    string `json:"remove"`
			Move      string `json:"move"`
		} `json:"links"`

		ACL struct {
			App            string `json:"app"`
			CanChangePerms bool   `json:"can_change_perms"`
			CanList        bool   `json:"can_list"`
			CanShareApp    bool   `json:"can_share_app"`
			CanShareGlobal bool   `json:"can_share_global"`
			CanShareUser   bool   `json:"can_share_user"`
			CanWrite       bool   `json:"can_write"`
			Modifiable     bool   `json:"modifiable"`
			Owner          string `json:"owner"`
			Perms          struct {
				Read  []string `json:"read"`
				Write []string `json:"write"`
			} `json:"perms"`
			Removable bool   `json:"removable"`
			Sharing   string `json:"sharing"`
		} `json:"acl"`
		Content struct {
			Disabled    bool        `json:"disabled"`
			EaiACL      interface{} `json:"eai:acl"`
			EaiAppName  string      `json:"eai:appName"`
			EaiData     string      `json:"eai:data"`
			EaiDigest   string      `json:"eai:digest"`
			EaiType     string      `json:"eai:type"`
			EaiUserName string      `json:"eai:userName"`
			IsDashboard bool        `json:"isDashboard"`
			IsVisible   bool        `json:"isVisible"`
			Label       string      `json:"label"`
			RootNode    string      `json:"rootNode"`
			Version     string      `json:"version"`
		} `json:"content"`
	} `json:"entry"`
	Paging struct {
		Total   int `json:"total"`
		PerPage int `json:"perPage"`
		Offset  int `json:"offset"`
	} `json:"paging"`
	Messages []interface{} `json:"messages"`
}

// SplunkReportDataUIView is what is returned from the "Report" view
type SplunkReportDataUIView struct {
	Links struct {
		Create string `json:"create"`
		Reload string `json:"_reload"`
		ACL    string `json:"_acl"`
	} `json:"links"`
	Origin    string `json:"origin"`
	Updated   string `json:"updated"`
	Generator struct {
		Build   string `json:"build"`
		Version string `json:"version"`
	} `json:"generator"`
	Entry []struct {
		Name    string `json:"name"`
		ID      string `json:"id"`
		Updated string `json:"updated"`
		Links   struct {
			Alternate string `json:"alternate"`
			List      string `json:"list"`
			Reload    string `json:"_reload"`
			Edit      string `json:"edit"`
			Remove    string `json:"remove"`
			Move      string `json:"move"`
			Disable   string `json:"disable"`
			Dispatch  string `json:"dispatch"`
			Embed     string `json:"embed"`
			History   string `json:"history"`
		} `json:"links"`
		Author string `json:"author"`
		ACL    struct {
			App            string      `json:"app"`
			CanChangePerms bool        `json:"can_change_perms"`
			CanList        bool        `json:"can_list"`
			CanShareApp    bool        `json:"can_share_app"`
			CanShareGlobal bool        `json:"can_share_global"`
			CanShareUser   bool        `json:"can_share_user"`
			CanWrite       bool        `json:"can_write"`
			Modifiable     bool        `json:"modifiable"`
			Owner          string      `json:"owner"`
			Perms          interface{} `json:"perms"`
			Removable      bool        `json:"removable"`
			Sharing        string      `json:"sharing"`
		} `json:"acl"`
		Content struct {
			ActionEmail                                                           bool        `json:"action.email"`
			ActionEmailSendresults                                                interface{} `json:"action.email.sendresults"`
			ActionEmailTo                                                         string      `json:"action.email.to"`
			ActionEmailUseNSSubject                                               string      `json:"action.email.useNSSubject"`
			ActionKeyindicatorInvert                                              string      `json:"action.keyindicator.invert"`
			ActionMakestreamsParamVerbose                                         string      `json:"action.makestreams.param.verbose"`
			ActionNbtstatParamVerbose                                             string      `json:"action.nbtstat.param.verbose"`
			ActionNotableParamVerbose                                             string      `json:"action.notable.param.verbose"`
			ActionNslookupParamVerbose                                            string      `json:"action.nslookup.param.verbose"`
			ActionPingParamVerbose                                                string      `json:"action.ping.param.verbose"`
			ActionPopulateLookup                                                  bool        `json:"action.populate_lookup"`
			ActionRiskParamVerbose                                                string      `json:"action.risk.param.verbose"`
			ActionRss                                                             bool        `json:"action.rss"`
			ActionScript                                                          bool        `json:"action.script"`
			ActionSend2UbaParamVerbose                                            string      `json:"action.send2uba.param.verbose"`
			ActionSummaryIndex                                                    bool        `json:"action.summary_index"`
			ActionThreatAddParamVerbose                                           string      `json:"action.threat_add.param.verbose"`
			Actions                                                               string      `json:"actions"`
			AlertDigestMode                                                       bool        `json:"alert.digest_mode"`
			AlertExpires                                                          string      `json:"alert.expires"`
			AlertManagedBy                                                        string      `json:"alert.managedBy"`
			AlertSeverity                                                         int         `json:"alert.severity"`
			AlertSuppress                                                         interface{} `json:"alert.suppress"`
			AlertSuppressFields                                                   string      `json:"alert.suppress.fields"`
			AlertSuppressPeriod                                                   string      `json:"alert.suppress.period"`
			AlertTrack                                                            bool        `json:"alert.track"`
			AlertComparator                                                       string      `json:"alert_comparator"`
			AlertCondition                                                        string      `json:"alert_condition"`
			AlertThreshold                                                        string      `json:"alert_threshold"`
			AlertType                                                             string      `json:"alert_type"`
			AllowSkew                                                             string      `json:"allow_skew"`
			AutoSummarize                                                         bool        `json:"auto_summarize"`
			AutoSummarizeCommand                                                  string      `json:"auto_summarize.command"`
			AutoSummarizeCronSchedule                                             string      `json:"auto_summarize.cron_schedule"`
			AutoSummarizeDispatchEarliestTime                                     string      `json:"auto_summarize.dispatch.earliest_time"`
			AutoSummarizeDispatchLatestTime                                       string      `json:"auto_summarize.dispatch.latest_time"`
			AutoSummarizeDispatchTimeFormat                                       string      `json:"auto_summarize.dispatch.time_format"`
			AutoSummarizeDispatchTTL                                              string      `json:"auto_summarize.dispatch.ttl"`
			AutoSummarizeMaxConcurrent                                            string      `json:"auto_summarize.max_concurrent"`
			AutoSummarizeMaxDisabledBuckets                                       int         `json:"auto_summarize.max_disabled_buckets"`
			AutoSummarizeMaxSummaryRatio                                          float64     `json:"auto_summarize.max_summary_ratio"`
			AutoSummarizeMaxSummarySize                                           int         `json:"auto_summarize.max_summary_size"`
			AutoSummarizeMaxTime                                                  string      `json:"auto_summarize.max_time"`
			AutoSummarizeSuspendPeriod                                            string      `json:"auto_summarize.suspend_period"`
			AutoSummarizeTimespan                                                 string      `json:"auto_summarize.timespan"`
			CronSchedule                                                          string      `json:"cron_schedule"`
			DeferScheduledSearchableIdxc                                          string      `json:"defer_scheduled_searchable_idxc"`
			Description                                                           string      `json:"description"`
			Disabled                                                              bool        `json:"disabled"`
			DispatchAutoCancel                                                    string      `json:"dispatch.auto_cancel"`
			DispatchAutoPause                                                     string      `json:"dispatch.auto_pause"`
			DispatchBuckets                                                       int         `json:"dispatch.buckets"`
			DispatchEarliestTime                                                  string      `json:"dispatch.earliest_time"`
			DispatchIndexEarliest                                                 string      `json:"dispatch.index_earliest"`
			DispatchIndexLatest                                                   string      `json:"dispatch.index_latest"`
			DispatchIndexedRealtime                                               interface{} `json:"dispatch.indexedRealtime"`
			DispatchIndexedRealtimeMinSpan                                        string      `json:"dispatch.indexedRealtimeMinSpan"`
			DispatchIndexedRealtimeOffset                                         string      `json:"dispatch.indexedRealtimeOffset"`
			DispatchLatestTime                                                    string      `json:"dispatch.latest_time"`
			DispatchLookups                                                       bool        `json:"dispatch.lookups"`
			DispatchMaxCount                                                      int         `json:"dispatch.max_count"`
			DispatchMaxTime                                                       int         `json:"dispatch.max_time"`
			DispatchReduceFreq                                                    int         `json:"dispatch.reduce_freq"`
			DispatchRtBackfill                                                    bool        `json:"dispatch.rt_backfill"`
			DispatchRtMaximumSpan                                                 string      `json:"dispatch.rt_maximum_span"`
			DispatchSampleRatio                                                   string      `json:"dispatch.sample_ratio"`
			DispatchSpawnProcess                                                  bool        `json:"dispatch.spawn_process"`
			DispatchTimeFormat                                                    string      `json:"dispatch.time_format"`
			DispatchTTL                                                           string      `json:"dispatch.ttl"`
			DispatchAs                                                            string      `json:"dispatchAs"`
			DisplayEventsFields                                                   string      `json:"display.events.fields"`
			DisplayEventsListDrilldown                                            string      `json:"display.events.list.drilldown"`
			DisplayEventsListWrap                                                 string      `json:"display.events.list.wrap"`
			DisplayEventsMaxLines                                                 string      `json:"display.events.maxLines"`
			DisplayEventsRawDrilldown                                             string      `json:"display.events.raw.drilldown"`
			DisplayEventsRowNumbers                                               string      `json:"display.events.rowNumbers"`
			DisplayEventsTableDrilldown                                           string      `json:"display.events.table.drilldown"`
			DisplayEventsTableWrap                                                string      `json:"display.events.table.wrap"`
			DisplayEventsType                                                     string      `json:"display.events.type"`
			DisplayGeneralEnablePreview                                           string      `json:"display.general.enablePreview"`
			DisplayGeneralMigratedFromViewState                                   string      `json:"display.general.migratedFromViewState"`
			DisplayGeneralTimeRangePickerShow                                     string      `json:"display.general.timeRangePicker.show"`
			DisplayGeneralType                                                    string      `json:"display.general.type"`
			DisplayPageSearchMode                                                 string      `json:"display.page.search.mode"`
			DisplayPageSearchPatternsSensitivity                                  string      `json:"display.page.search.patterns.sensitivity"`
			DisplayPageSearchShowFields                                           string      `json:"display.page.search.showFields"`
			DisplayPageSearchTab                                                  string      `json:"display.page.search.tab"`
			DisplayPageSearchTimelineFormat                                       string      `json:"display.page.search.timeline.format"`
			DisplayPageSearchTimelineScale                                        string      `json:"display.page.search.timeline.scale"`
			DisplayStatisticsDrilldown                                            string      `json:"display.statistics.drilldown"`
			DisplayStatisticsOverlay                                              string      `json:"display.statistics.overlay"`
			DisplayStatisticsPercentagesRow                                       string      `json:"display.statistics.percentagesRow"`
			DisplayStatisticsRowNumbers                                           string      `json:"display.statistics.rowNumbers"`
			DisplayStatisticsShow                                                 string      `json:"display.statistics.show"`
			DisplayStatisticsTotalsRow                                            string      `json:"display.statistics.totalsRow"`
			DisplayStatisticsWrap                                                 string      `json:"display.statistics.wrap"`
			DisplayVisualizationsChartHeight                                      string      `json:"display.visualizations.chartHeight"`
			DisplayVisualizationsChartingAxisLabelsXMajorLabelStyleOverflowMode   string      `json:"display.visualizations.charting.axisLabelsX.majorLabelStyle.overflowMode"`
			DisplayVisualizationsChartingAxisLabelsXMajorLabelStyleRotation       string      `json:"display.visualizations.charting.axisLabelsX.majorLabelStyle.rotation"`
			DisplayVisualizationsChartingAxisLabelsXMajorUnit                     string      `json:"display.visualizations.charting.axisLabelsX.majorUnit"`
			DisplayVisualizationsChartingAxisLabelsYMajorUnit                     string      `json:"display.visualizations.charting.axisLabelsY.majorUnit"`
			DisplayVisualizationsChartingAxisLabelsY2MajorUnit                    string      `json:"display.visualizations.charting.axisLabelsY2.majorUnit"`
			DisplayVisualizationsChartingAxisTitleXText                           string      `json:"display.visualizations.charting.axisTitleX.text"`
			DisplayVisualizationsChartingAxisTitleXVisibility                     string      `json:"display.visualizations.charting.axisTitleX.visibility"`
			DisplayVisualizationsChartingAxisTitleYText                           string      `json:"display.visualizations.charting.axisTitleY.text"`
			DisplayVisualizationsChartingAxisTitleYVisibility                     string      `json:"display.visualizations.charting.axisTitleY.visibility"`
			DisplayVisualizationsChartingAxisTitleY2Text                          string      `json:"display.visualizations.charting.axisTitleY2.text"`
			DisplayVisualizationsChartingAxisTitleY2Visibility                    string      `json:"display.visualizations.charting.axisTitleY2.visibility"`
			DisplayVisualizationsChartingAxisXAbbreviation                        string      `json:"display.visualizations.charting.axisX.abbreviation"`
			DisplayVisualizationsChartingAxisXMaximumNumber                       string      `json:"display.visualizations.charting.axisX.maximumNumber"`
			DisplayVisualizationsChartingAxisXMinimumNumber                       string      `json:"display.visualizations.charting.axisX.minimumNumber"`
			DisplayVisualizationsChartingAxisXScale                               string      `json:"display.visualizations.charting.axisX.scale"`
			DisplayVisualizationsChartingAxisYAbbreviation                        string      `json:"display.visualizations.charting.axisY.abbreviation"`
			DisplayVisualizationsChartingAxisYMaximumNumber                       string      `json:"display.visualizations.charting.axisY.maximumNumber"`
			DisplayVisualizationsChartingAxisYMinimumNumber                       string      `json:"display.visualizations.charting.axisY.minimumNumber"`
			DisplayVisualizationsChartingAxisYScale                               string      `json:"display.visualizations.charting.axisY.scale"`
			DisplayVisualizationsChartingAxisY2Abbreviation                       string      `json:"display.visualizations.charting.axisY2.abbreviation"`
			DisplayVisualizationsChartingAxisY2Enabled                            string      `json:"display.visualizations.charting.axisY2.enabled"`
			DisplayVisualizationsChartingAxisY2MaximumNumber                      string      `json:"display.visualizations.charting.axisY2.maximumNumber"`
			DisplayVisualizationsChartingAxisY2MinimumNumber                      string      `json:"display.visualizations.charting.axisY2.minimumNumber"`
			DisplayVisualizationsChartingAxisY2Scale                              string      `json:"display.visualizations.charting.axisY2.scale"`
			DisplayVisualizationsChartingChart                                    string      `json:"display.visualizations.charting.chart"`
			DisplayVisualizationsChartingChartBubbleMaximumSize                   string      `json:"display.visualizations.charting.chart.bubbleMaximumSize"`
			DisplayVisualizationsChartingChartBubbleMinimumSize                   string      `json:"display.visualizations.charting.chart.bubbleMinimumSize"`
			DisplayVisualizationsChartingChartBubbleSizeBy                        string      `json:"display.visualizations.charting.chart.bubbleSizeBy"`
			DisplayVisualizationsChartingChartNullValueMode                       string      `json:"display.visualizations.charting.chart.nullValueMode"`
			DisplayVisualizationsChartingChartOverlayFields                       string      `json:"display.visualizations.charting.chart.overlayFields"`
			DisplayVisualizationsChartingChartRangeValues                         string      `json:"display.visualizations.charting.chart.rangeValues"`
			DisplayVisualizationsChartingChartShowDataLabels                      string      `json:"display.visualizations.charting.chart.showDataLabels"`
			DisplayVisualizationsChartingChartSliceCollapsingThreshold            string      `json:"display.visualizations.charting.chart.sliceCollapsingThreshold"`
			DisplayVisualizationsChartingChartStackMode                           string      `json:"display.visualizations.charting.chart.stackMode"`
			DisplayVisualizationsChartingChartStyle                               string      `json:"display.visualizations.charting.chart.style"`
			DisplayVisualizationsChartingDrilldown                                string      `json:"display.visualizations.charting.drilldown"`
			DisplayVisualizationsChartingFieldDashStyles                          string      `json:"display.visualizations.charting.fieldDashStyles"`
			DisplayVisualizationsChartingGaugeColors                              string      `json:"display.visualizations.charting.gaugeColors"`
			DisplayVisualizationsChartingLayoutSplitSeries                        string      `json:"display.visualizations.charting.layout.splitSeries"`
			DisplayVisualizationsChartingLayoutSplitSeriesAllowIndependentYRanges string      `json:"display.visualizations.charting.layout.splitSeries.allowIndependentYRanges"`
			DisplayVisualizationsChartingLegendLabelStyleOverflowMode             string      `json:"display.visualizations.charting.legend.labelStyle.overflowMode"`
			DisplayVisualizationsChartingLegendMode                               string      `json:"display.visualizations.charting.legend.mode"`
			DisplayVisualizationsChartingLegendPlacement                          string      `json:"display.visualizations.charting.legend.placement"`
			DisplayVisualizationsChartingLineWidth                                string      `json:"display.visualizations.charting.lineWidth"`
			DisplayVisualizationsCustomDrilldown                                  string      `json:"display.visualizations.custom.drilldown"`
			DisplayVisualizationsCustomHeight                                     string      `json:"display.visualizations.custom.height"`
			DisplayVisualizationsCustomType                                       string      `json:"display.visualizations.custom.type"`
			DisplayVisualizationsMapHeight                                        string      `json:"display.visualizations.mapHeight"`
			DisplayVisualizationsMappingChoroplethLayerColorBins                  string      `json:"display.visualizations.mapping.choroplethLayer.colorBins"`
			DisplayVisualizationsMappingChoroplethLayerColorMode                  string      `json:"display.visualizations.mapping.choroplethLayer.colorMode"`
			DisplayVisualizationsMappingChoroplethLayerMaximumColor               string      `json:"display.visualizations.mapping.choroplethLayer.maximumColor"`
			DisplayVisualizationsMappingChoroplethLayerMinimumColor               string      `json:"display.visualizations.mapping.choroplethLayer.minimumColor"`
			DisplayVisualizationsMappingChoroplethLayerNeutralPoint               string      `json:"display.visualizations.mapping.choroplethLayer.neutralPoint"`
			DisplayVisualizationsMappingChoroplethLayerShapeOpacity               string      `json:"display.visualizations.mapping.choroplethLayer.shapeOpacity"`
			DisplayVisualizationsMappingChoroplethLayerShowBorder                 string      `json:"display.visualizations.mapping.choroplethLayer.showBorder"`
			DisplayVisualizationsMappingDataMaxClusters                           string      `json:"display.visualizations.mapping.data.maxClusters"`
			DisplayVisualizationsMappingDrilldown                                 string      `json:"display.visualizations.mapping.drilldown"`
			DisplayVisualizationsMappingLegendPlacement                           string      `json:"display.visualizations.mapping.legend.placement"`
			DisplayVisualizationsMappingMapCenter                                 string      `json:"display.visualizations.mapping.map.center"`
			DisplayVisualizationsMappingMapPanning                                string      `json:"display.visualizations.mapping.map.panning"`
			DisplayVisualizationsMappingMapScrollZoom                             string      `json:"display.visualizations.mapping.map.scrollZoom"`
			DisplayVisualizationsMappingMapZoom                                   string      `json:"display.visualizations.mapping.map.zoom"`
			DisplayVisualizationsMappingMarkerLayerMarkerMaxSize                  string      `json:"display.visualizations.mapping.markerLayer.markerMaxSize"`
			DisplayVisualizationsMappingMarkerLayerMarkerMinSize                  string      `json:"display.visualizations.mapping.markerLayer.markerMinSize"`
			DisplayVisualizationsMappingMarkerLayerMarkerOpacity                  string      `json:"display.visualizations.mapping.markerLayer.markerOpacity"`
			DisplayVisualizationsMappingShowTiles                                 string      `json:"display.visualizations.mapping.showTiles"`
			DisplayVisualizationsMappingTileLayerMaxZoom                          string      `json:"display.visualizations.mapping.tileLayer.maxZoom"`
			DisplayVisualizationsMappingTileLayerMinZoom                          string      `json:"display.visualizations.mapping.tileLayer.minZoom"`
			DisplayVisualizationsMappingTileLayerTileOpacity                      string      `json:"display.visualizations.mapping.tileLayer.tileOpacity"`
			DisplayVisualizationsMappingTileLayerURL                              string      `json:"display.visualizations.mapping.tileLayer.url"`
			DisplayVisualizationsMappingType                                      string      `json:"display.visualizations.mapping.type"`
			DisplayVisualizationsShow                                             string      `json:"display.visualizations.show"`
			DisplayVisualizationsSinglevalueAfterLabel                            string      `json:"display.visualizations.singlevalue.afterLabel"`
			DisplayVisualizationsSinglevalueBeforeLabel                           string      `json:"display.visualizations.singlevalue.beforeLabel"`
			DisplayVisualizationsSinglevalueColorBy                               string      `json:"display.visualizations.singlevalue.colorBy"`
			DisplayVisualizationsSinglevalueColorMode                             string      `json:"display.visualizations.singlevalue.colorMode"`
			DisplayVisualizationsSinglevalueDrilldown                             string      `json:"display.visualizations.singlevalue.drilldown"`
			DisplayVisualizationsSinglevalueNumberPrecision                       string      `json:"display.visualizations.singlevalue.numberPrecision"`
			DisplayVisualizationsSinglevalueRangeColors                           string      `json:"display.visualizations.singlevalue.rangeColors"`
			DisplayVisualizationsSinglevalueRangeValues                           string      `json:"display.visualizations.singlevalue.rangeValues"`
			DisplayVisualizationsSinglevalueShowSparkline                         string      `json:"display.visualizations.singlevalue.showSparkline"`
			DisplayVisualizationsSinglevalueShowTrendIndicator                    string      `json:"display.visualizations.singlevalue.showTrendIndicator"`
			DisplayVisualizationsSinglevalueTrendColorInterpretation              string      `json:"display.visualizations.singlevalue.trendColorInterpretation"`
			DisplayVisualizationsSinglevalueTrendDisplayMode                      string      `json:"display.visualizations.singlevalue.trendDisplayMode"`
			DisplayVisualizationsSinglevalueTrendInterval                         string      `json:"display.visualizations.singlevalue.trendInterval"`
			DisplayVisualizationsSinglevalueUnderLabel                            string      `json:"display.visualizations.singlevalue.underLabel"`
			DisplayVisualizationsSinglevalueUnit                                  string      `json:"display.visualizations.singlevalue.unit"`
			DisplayVisualizationsSinglevalueUnitPosition                          string      `json:"display.visualizations.singlevalue.unitPosition"`
			DisplayVisualizationsSinglevalueUseColors                             string      `json:"display.visualizations.singlevalue.useColors"`
			DisplayVisualizationsSinglevalueUseThousandSeparators                 string      `json:"display.visualizations.singlevalue.useThousandSeparators"`
			DisplayVisualizationsSinglevalueHeight                                string      `json:"display.visualizations.singlevalueHeight"`
			DisplayVisualizationsTrellisEnabled                                   string      `json:"display.visualizations.trellis.enabled"`
			DisplayVisualizationsTrellisScalesShared                              string      `json:"display.visualizations.trellis.scales.shared"`
			DisplayVisualizationsTrellisSize                                      string      `json:"display.visualizations.trellis.size"`
			DisplayVisualizationsTrellisSplitBy                                   string      `json:"display.visualizations.trellis.splitBy"`
			DisplayVisualizationsType                                             string      `json:"display.visualizations.type"`
			Displayview                                                           string      `json:"displayview"`
			EaiACL                                                                interface{} `json:"eai:acl"`
			EmbedEnabled                                                          string      `json:"embed.enabled"`
			IsScheduled                                                           bool        `json:"is_scheduled"`
			IsVisible                                                             bool        `json:"is_visible"`
			MaxConcurrent                                                         int         `json:"max_concurrent"`
			NextScheduledTime                                                     string      `json:"next_scheduled_time"`
			QualifiedSearch                                                       string      `json:"qualifiedSearch"`
			RealtimeSchedule                                                      bool        `json:"realtime_schedule"`
			RequestUIDispatchApp                                                  string      `json:"request.ui_dispatch_app"`
			RequestUIDispatchView                                                 string      `json:"request.ui_dispatch_view"`
			RestartOnSearchpeerAdd                                                bool        `json:"restart_on_searchpeer_add"`
			RunNTimes                                                             int         `json:"run_n_times"`
			RunOnStartup                                                          bool        `json:"run_on_startup"`
			SchedulePriority                                                      string      `json:"schedule_priority"`
			ScheduleWindow                                                        string      `json:"schedule_window"`
			Search                                                                string      `json:"search"`
			Vsid                                                                  string      `json:"vsid"`
		} `json:"content"`
	} `json:"entry"`
	Paging struct {
		Total   int `json:"total"`
		PerPage int `json:"perPage"`
		Offset  int `json:"offset"`
	} `json:"paging"`
	Messages []interface{} `json:"messages"`
}

// SplunkJobSubmission after a search job is subbmited, the sid is for tracking it.
type SplunkJobSubmission struct {
	SID string `json:"sid"`
}

// SplunkSearchStatusUIView gives the details of a search job that was submitted
type SplunkSearchStatusUIView struct {
	Links struct {
	} `json:"links"`
	Origin    string `json:"origin"`
	Updated   string `json:"updated"`
	Generator struct {
		Build   string `json:"build"`
		Version string `json:"version"`
	} `json:"generator"`
	Entry []struct {
		Name    string `json:"name"`
		ID      string `json:"id"`
		Updated string `json:"updated"`
		Links   struct {
			Alternate      string `json:"alternate"`
			Events         string `json:"events"`
			Results        string `json:"results"`
			ResultsPreview string `json:"results_preview"`
			Timeline       string `json:"timeline"`
			Summary        string `json:"summary"`
			Control        string `json:"control"`
		} `json:"links"`
		Published string `json:"published"`
		Author    string `json:"author"`
		Content   struct {
			BundleVersion  string        `json:"bundleVersion"`
			CursorTime     string        `json:"cursorTime"`
			DefaultSaveTTL string        `json:"defaultSaveTTL"`
			DefaultTTL     string        `json:"defaultTTL"`
			Delegate       string        `json:"delegate"`
			DiskUsage      int           `json:"diskUsage"`
			DispatchState  string        `json:"dispatchState"`
			DoneProgress   int           `json:"doneProgress"`
			EarliestTime   string        `json:"earliestTime"`
			IsDone         bool          `json:"isDone"`
			IsFailed       bool          `json:"isFailed"`
			IsFinalized    bool          `json:"isFinalized"`
			IsPaused       bool          `json:"isPaused"`
			IsSaved        bool          `json:"isSaved"`
			IsSavedSearch  bool          `json:"isSavedSearch"`
			IsZombie       bool          `json:"isZombie"`
			Label          string        `json:"label"`
			Sid            string        `json:"sid"`
			StatusBuckets  int           `json:"statusBuckets"`
			TTL            int           `json:"ttl"`
			Messages       []interface{} `json:"messages"`
			Request        struct {
				AutoCancel    string `json:"auto_cancel"`
				EarliestTime  string `json:"earliest_time"`
				Preview       string `json:"preview"`
				Provenance    string `json:"provenance"`
				Rf            string `json:"rf"`
				Search        string `json:"search"`
				StatusBuckets string `json:"status_buckets"`
			} `json:"request"`
			Runtime struct {
				AutoCancel string `json:"auto_cancel"`
				AutoPause  string `json:"auto_pause"`
			} `json:"runtime"`
			SearchProviders []interface{} `json:"searchProviders"`
		} `json:"content"`
		ACL struct {
			Perms struct {
				Read  []string `json:"read"`
				Write []string `json:"write"`
			} `json:"perms"`
			Owner      string `json:"owner"`
			Modifiable bool   `json:"modifiable"`
			Sharing    string `json:"sharing"`
			App        string `json:"app"`
			CanWrite   bool   `json:"can_write"`
			TTL        string `json:"ttl"`
		} `json:"acl"`
	} `json:"entry"`
	Paging struct {
		Total   int `json:"total"`
		PerPage int `json:"perPage"`
		Offset  int `json:"offset"`
	} `json:"paging"`
}

// SplunkSearchResultsUIView is returned for each page of Search Results
type SplunkSearchResultsUIView struct {
	Preview          bool `json:"preview"`
	InitOffset       int  `json:"init_offset"`
	PostProcessCount int  `json:"post_process_count"`
	Messages         []struct {
		Type string `json:"type"`
		Text string `json:"text"`
	} `json:"messages"`
	Fields []struct {
		Name string `json:"name"`
	} `json:"fields"`
	Results []struct {
		Time         string `json:"_time"`
		APIEt        string `json:"api_et,omitempty"`
		APILt        string `json:"api_lt,omitempty"`
		EventCount   string `json:"event_count"`
		ExecTime     string `json:"exec_time"`
		IsRealtime   string `json:"is_realtime"`
		Provenance   string `json:"provenance,omitempty"`
		ResultCount  string `json:"result_count"`
		ScanCount    string `json:"scan_count"`
		Search       string `json:"search"`
		SearchEt     string `json:"search_et,omitempty"`
		SearchLt     string `json:"search_lt,omitempty"`
		Sid          string `json:"sid"`
		SplunkServer string `json:"splunk_server"`
		Status       string `json:"status"`
		TotalRunTime string `json:"total_run_time"`
	} `json:"results"`
	Highlighted struct {
	} `json:"highlighted"`
}
