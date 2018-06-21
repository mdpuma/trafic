#!/bin/bash

set -eu

function mklabel() {
	local exid=$1
	local unixtime=$(date +%s)

	printf "lola-%s-%s" "${exid}" "${unixtime}"
}

# remote measurements don't include capturing traffic
#IFACE=${IFACE:-eth0}
IFACE=
EXID=${EXID:-baseline}
HOST=${HOST:-iperf}

for load in 75 80 85 90 95
do
	exid="${EXID}-${load}"
	label=$(mklabel "${exid}")
	capfn="${label}.pcap"

	# start servers
	wget --header "X-CONF: ${exid}.env" \
		-O /dev/null \
		http://${HOST}-server:9000/hooks/start-servers

	sleep 1

	# start clients
	wget --header "X-CONF: ${exid}.env" \
		--header "X-LABEL: ${label}" \
		--header "X-DB: ${EXID}" \
		-O /dev/null \
		http://${HOST}-client:9000/hooks/start-clients

	if [ -n "${IFACE}" ]; then
		# start capture for 60s
		tshark -i ${IFACE} -s 128 -w ${capfn} -f 'tcp or udp' -a duration:60
		# try to save as much space as possible
		bzip2 -9 ${capfn}
		sleep 5	# allow some time for flows to drain
	else
		sleep 65
	fi
	# cleanup (and, possibly, go again)
	wget http://${HOST}-server:9000/hooks/stop-servers -O /dev/null
	wget http://${HOST}-client:9000/hooks/stop-clients -O /dev/null
	sleep 5
done
