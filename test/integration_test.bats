#!/usr/bin/env bats

load bats_utils

setup(){

 #   if [ "$BATS_TEST_NUMBER" -eq "1" ]; then
 #       teardown
 #   fi

    run kubectl scale --replicas=1 deployment/kubernetes-grafana-test
    run kubectl scale --replicas=1 deployment/grafana

    validateGrafanaUrl
}

teardown(){
    dumpState

    kubectl delete events --all

    run kubectl scale --replicas=0 deployment/kubernetes-grafana-test
    run kubectl scale --replicas=0 deployment/grafana

    kubectl delete Dashboard --ignore-not-found=true --all
    kubectl delete NotificationChannel --ignore-not-found=true --all
    kubectl delete DataSource --ignore-not-found=true --all

    # clean up comparison files if they exist
    rm -f a.json
    rm -f b.json
}

#
# dashboards
#
@test "creating a Dashboard object creates a Grafana Dashboard" {
    count=0

    for filename in dashboards/*.yaml; do
        dashboardId=$(validatePostDashboard $filename)

        echo "Test Creating $filename ($dashboardId)"

        (( count++ ))
        validateDashboardCount $count

        validateEventCount Dashboard Synced $(objectNameFromFile $filename) 1
    done
}

@test "deleting a Dashboard object deletes the Grafana Dashboard" {
    for filename in dashboards/*.yaml; do
        dashboardId=$(validatePostDashboard $filename)

        echo "Test Deleting $filename ($dashboardId)"

        kubectl delete -f $filename

        sleep 5s

        httpStatus=$(curl --silent --output /dev/null --write-out "%{http_code}" ${GRAFANA_URL}/api/dashboards/uid/${dashboardId})

        [ "$httpStatus" -eq "404" ]

        validateDashboardCount 0

        validateEventCount Dashboard Synced $(objectNameFromFile $filename) 1
        validateEventCount Dashboard Deleted $(objectNameFromFile $filename) 1
    done
}

@test "deleting a Dashboard while the controller is not running deletes the dashboard in Grafana" {
    for filename in dashboards/*.yaml; do
        dashboardId=$(validatePostDashboard $filename)

        kubectl scale --replicas=0 deployment/kubernetes-grafana-test

        echo "Test Deleting $filename ($dashboardId)"

        kubectl delete -f $filename

        sleep 5s

        kubectl scale --replicas=1 deployment/kubernetes-grafana-test

        sleep 10s

        httpStatus=$(curl --silent --output /dev/null --write-out "%{http_code}" ${GRAFANA_URL}/api/dashboards/uid/${dashboardId})

        [ "$httpStatus" -eq "404" ]

        validateDashboardCount 0

        validateEventCount Dashboard Synced $(objectNameFromFile $filename) 1
    done
}

@test "creating a Dashboard object creates the same dashboard in Grafana" {
    count=0

    for filename in dashboards/*.yaml; do
        validateDashboardContents $filename

        (( count++ ))
        validateDashboardCount $count

        validateEventCount Dashboard Synced $(objectNameFromFile $filename) 1
    done
}

@test "updating a Dashboard object updates the dashboard in Grafana" {
    count=0
    
    for filename in dashboards/*.yaml; do
        validateDashboardContents $filename

        (( count++ ))
        validateDashboardCount $count

        validateEventCount Dashboard Synced $(objectNameFromFile $filename) 1
    done

    # the .update files have dashboards with the same ids and different contents. 
    #  not the best.  not the worst.  could be improved.
    for filename in dashboards/*.update; do
        validateDashboardContents $filename

        validateDashboardCount $count

        validateEventCount Dashboard Synced $(objectNameFromFile $filename) 2
    done
}

#
# notification channels
#
@test "creating a NotificationChannel object creates a Grafana Notification Channel" {
    for filename in notification_channels/*.yaml; do
        channelId=$(validatePostNotificationChannel $filename)

        echo "Test Creating $filename ($channelId)"

        (( count++ ))
        validateNotificationChannelCount $count

        validateEventCount NotificationChannel Synced $(objectNameFromFile $filename) 1
    done
}

@test "deleting a NotificationChannel object deletes the Grafana Notification Channel" {

    for filename in notification_channels/*.yaml; do
        channelId=$(validatePostNotificationChannel $filename)

        echo "Test Deleting $filename ($channelId)"

        kubectl delete -f $filename

        sleep 5s

        httpStatus=$(curl --silent --output /dev/null --write-out "%{http_code}" ${GRAFANA_URL}/api/alert-notifications/${channelId})

        # for some reason grafana 500s when you GET a non-existent alert-notifications?
        [ "$httpStatus" -eq "500" ]

        validateNotificationChannelCount 0

        validateEventCount NotificationChannel Synced $(objectNameFromFile $filename) 1
        validateEventCount NotificationChannel Deleted $(objectNameFromFile $filename) 1
    done
}

@test "deleting a NotificationChannel while the controller is not running deletes the notification channel in Grafana" {

    for filename in notification_channels/*.yaml; do
        channelId=$(validatePostNotificationChannel $filename)

        kubectl scale --replicas=0 deployment/kubernetes-grafana-test

        echo "Test Deleting $filename ($channelId)"

        kubectl delete -f $filename

        sleep 5s

        kubectl scale --replicas=1 deployment/kubernetes-grafana-test

        sleep 10s

        httpStatus=$(curl --silent --output /dev/null --write-out "%{http_code}" ${GRAFANA_URL}/api/alert-notifications/${channelId})

        # for some reason grafana 500s when you GET a non-existent alert-notifications?
        [ "$httpStatus" -eq "500" ]

        validateNotificationChannelCount 0

        validateEventCount NotificationChannel Synced $(objectNameFromFile $filename) 1
    done
}


@test "creating a NotificationChannel object creates the same channel in Grafana" {
    count=0

    for filename in notification_channels/*.yaml; do
        validateNotificationChannelContents $filename

        (( count++ ))
        validateNotificationChannelCount $count

        validateEventCount NotificationChannel Synced $(objectNameFromFile $filename) 1
    done
}

@test "updating a NotificationChannel object updates the channel in Grafana" {
    count=0
    
    for filename in notification_channels/*.yaml; do
        validateNotificationChannelContents $filename

        (( count++ ))
        validateNotificationChannelCount $count

        validateEventCount NotificationChannel Synced $(objectNameFromFile $filename) 1
    done

    for filename in notification_channels/*.update; do
        validateNotificationChannelContents $filename

        validateNotificationChannelCount $count

        validateEventCount NotificationChannel Synced $(objectNameFromFile $filename) 2
    done
}

#
# data sources
#
@test "creating a DataSource object creates a Grafana DataSource" {
    count=0

    for filename in datasources/*.yaml; do
        sourceId=$(validatePostDataSource $filename)

        echo "Test Creating $filename ($sourceId)"

        (( count++ ))
        validateDataSourceCount $count

        validateEventCount DataSource Synced $(objectNameFromFile $filename) 1
    done
}

@test "deleting a DataSource object deletes the Grafana DataSource" {

    for filename in datasources/*.yaml; do
        sourceId=$(validatePostDataSource $filename)

        echo "Test Deleting $filename ($sourceId)"

        kubectl delete -f $filename

        sleep 5s

        httpStatus=$(curl --silent --output /dev/null --write-out "%{http_code}" ${GRAFANA_URL}/api/datasources/${sourceId})

        echo "status $httpStatus"
        curl ${GRAFANA_URL}/api/datasources

        [ "$httpStatus" -eq "404" ]

        validateDataSourceCount 0

        validateEventCount DataSource Synced $(objectNameFromFile $filename) 1
        validateEventCount DataSource Deleted $(objectNameFromFile $filename) 1
    done
}

@test "deleting a DataSource while the controller is not running deletes the datasource in Grafana" {

    for filename in datasources/*.yaml; do
        sourceId=$(validatePostDataSource $filename)

        kubectl scale --replicas=0 deployment/kubernetes-grafana-test

        echo "Test Deleting $filename ($sourceId)"

        kubectl delete -f $filename

        sleep 5s

        kubectl scale --replicas=1 deployment/kubernetes-grafana-test

        sleep 10s

        httpStatus=$(curl --silent --output /dev/null --write-out "%{http_code}" ${GRAFANA_URL}/api/datasources/${sourceId})

        echo "status $httpStatus"
        curl ${GRAFANA_URL}/api/datasources

        [ "$httpStatus" -eq "404" ]

        validateDataSourceCount 0

        validateEventCount DataSource Synced $(objectNameFromFile $filename) 1
    done
}

@test "creating a DataSource object creates the same datasource in Grafana" {
    count=0

    for filename in datasources/*.yaml; do
        validateDataSourceContents $filename

        (( count++ ))
        validateDataSourceCount $count

        validateEventCount DataSource Synced $(objectNameFromFile $filename) 1
    done
}

@test "updating a DataSource object updates the datasource in Grafana" {
    count=0
    
    for filename in datasources/*.yaml; do
        validateDataSourceContents $filename

        (( count++ ))
        validateDataSourceCount $count

        validateEventCount DataSource Synced $(objectNameFromFile $filename) 1
    done

    for filename in datasources/*.update; do
        validateDataSourceContents $filename

        validateDataSourceCount $count

        validateEventCount DataSource Synced $(objectNameFromFile $filename) 2
    done
}