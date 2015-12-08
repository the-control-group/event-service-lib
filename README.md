# Events Service Lib

## Statsd
All events should be prefixed with $service-name.$host-name.

 - Emit received events as `event.$event-name`
 - Emit invalid events as `invalid.$event-name`
 - Emit failed actions as `$action.failure.$event-name`
 - Emit successful actions as `$action.success.$event-name`
 - Emit stop signals as `stop.$signal`
 - Emit reloads as `reload`

In any case, use '_unknown' for `$event-name` if necessary